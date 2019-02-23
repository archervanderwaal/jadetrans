package engine

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"unicode/utf8"

	"github.com/archervanderwaal/jadetrans/config"
	"github.com/archervanderwaal/jadetrans/utils"
	"github.com/aybabtme/rgbterm"
)

var (
	errCodeMsgMap = map[string]string{
		"101": "缺少必填的参数，出现这个情况还可能是et的值和实际加密方式不对应",
		"102": "不支持的语言类型",
		"103": "翻译文本过长",
		"104": "不支持的API类型",
		"105": "不支持的签名类型",
		"106": "不支持的响应类型",
		"107": "不支持的传输加密类型",
		"108": "appKey无效,后台创建应用和实例并完成绑定,可获得应用ID和密钥等信息，其中应用ID就是appKey（注意不是应用密钥）",
		"109": "batchLog格式不正确",
		"110": "无相关服务的有效实例",
		"111": "开发者账号无效",
		"113": "q不能为空",
		"201": "解密失败，可能为DES,BASE64,URLDecode的错误",
		"202": "签名检验失败",
		"203": "访问IP地址不在可访问IP列表",
		"205": "请求的接口与应用的平台类型不一致，如有疑问请参考[入门指南]()",
		"206": "因为时间戳无效导致签名校验失败",
		"207": "重放请求",
		"301": "辞典查询失败",
		"302": "翻译查询失败",
		"303": "服务端的其它异常",
		"401": "账户已经欠费停",
		"411": "访问频率受限,请稍后访问",
		"412": "长请求过于频繁，请稍后访问",
	}
)

const (
	api            = "http://openapi.youdao.com/api"
	successErrCode = "0"
)

// YoudaoEngine contains some parameters that you want to use youdao Translation Service.
type YoudaoEngine struct {
	appKey    string
	appSecret string
	query     string
	from      string
	to        string
	voice     string
	signType  string
	curTime   string
	ext       string
	sign      string
	salt      string
}

// NewYoudaoEngine create a translation engine for you based on youdao Translation Service.
func NewYoudaoEngine(query, from, to, voice string, conf *config.Config) *YoudaoEngine {
	e := &YoudaoEngine{
		appKey:    conf.Youdao.AppKey,
		appSecret: conf.Youdao.AppSecret,
		query:     query,
		from:      from,
		to:        to,
		signType:  "v3",
		curTime:   utils.UTCTimestamp(),
		ext:       "mp3",
		salt:      utils.UUID(),
	}
	if voice == "0" || voice == "1" {
		e.voice = voice
	}
	sum := sha256.Sum256([]byte(fmt.Sprintf("%s%s%s%s%s", e.appKey,
		truncate(e.query), e.salt, e.curTime, e.appSecret)))
	e.sign = strings.ToLower(fmt.Sprintf("%x", sum))
	return e
}

// Query returns the translation results without errors.
func (e *YoudaoEngine) Query() (res string, err error) {
	var resp *http.Response
	spinnerID := utils.NewDefaultSpinnerAndStart("Querying...")
	defer utils.StopSpinner(spinnerID)
	resp, err = http.Get(e.requestURL())
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var bytes []byte
	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	result := e.generateResult(bytes)
	if err = result.success(); err != nil {
		return
	}
	res, err = result.format(), nil
	return
}

func (e *YoudaoEngine) requestURL() string {
	values := &url.Values{}
	values.Set("q", e.query)
	values.Set("from", e.from)
	values.Set("to", e.to)
	values.Set("appKey", e.appKey)
	values.Set("salt", e.salt)
	values.Set("sign", e.sign)
	values.Set("ext", e.ext)
	if e.voice != "" {
		values.Set("voice", e.voice)
	}
	values.Set("signType", e.signType)
	values.Set("curtime", e.curTime)
	return fmt.Sprintf("%s?%s", api, values.Encode())
}

func truncate(query string) string {
	queryRune := []rune(query)
	if queryLen := utf8.RuneCountInString(query); queryLen > 20 {
		return fmt.Sprintf("%s%d%s", string(queryRune[:10]), queryLen, string(queryRune[queryLen-10:]))
	}
	return query
}

func (e *YoudaoEngine) generateResult(jsonByteRes []byte) *result {
	res := &result{}
	json.Unmarshal(jsonByteRes, res)
	return res
}

type result struct {
	ErrorCode         string   `json:"errorCode"`
	Query             string   `json:"query"`
	SpeakURL          string   `json:"speakUrl"`
	TranslateSpeakURL string   `json:"tSpeakUrl"`
	Translation       []string `json:"translation"`
	Basic             basic    `json:"basic"`
	Web               []web    `json:"web"`
}

type basic struct {
	Phonetic   string   `json:"phonetic"`
	UkPhonetic string   `json:"uk-phonetic"`
	UsPhonetic string   `json:"us-phonetic"`
	Explains   []string `json:"explains"`
}

type web struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}

func (res *result) format() string {
	content := ""
	line := rgbterm.FgString("============================================================\n",
		0, 255, 0)

	content += line
	if res.Basic.UkPhonetic == "" && res.Basic.UsPhonetic == "" && res.Basic.Phonetic != "" {
		content += rgbterm.FgString(fmt.Sprintf("    %s: [%s]\n\n", "拼音", res.Basic.Phonetic),
			0, 255, 0)
	} else if res.Basic.UkPhonetic != "" || res.Basic.UsPhonetic != "" {
		content += rgbterm.FgString(fmt.Sprintf("    英: [%s]", res.Basic.UkPhonetic),
			0, 255, 0)
		content += rgbterm.FgString(fmt.Sprintf("    美: [%s]", res.Basic.UsPhonetic),
			0, 255, 0)
		content += "\n\n"
	}

	if len(res.Basic.Explains) == 0 {
		for _, tran := range res.Translation {
			content += rgbterm.FgString(fmt.Sprintf("    %s\n", tran), 0, 255, 0)
		}
	}

	if len(res.Basic.Explains) != 0 {
		for _, exp := range res.Basic.Explains {
			content += rgbterm.FgString(fmt.Sprintf("    %s", exp), 0, 255, 0)
			content += "\n"
		}
		content += "\n"
	}

	if len(res.Web) != 0 {
		for inx, w := range res.Web {
			content += rgbterm.FgString(fmt.Sprintf("    %d. %s\n", inx+1, w.Key), 0, 255, 0)
			content += rgbterm.FgString(fmt.Sprintf("       %s\n", strings.Join(w.Value, ", ")), 255, 0, 255)
		}
	}
	content += line
	return content
}

func (res *result) success() error {
	if res == nil || res.ErrorCode != successErrCode {
		errMsg := "未知的错误"
		if _, ok := errCodeMsgMap[res.ErrorCode]; ok {
			errMsg = errCodeMsgMap[res.ErrorCode]
		}
		return fmt.Errorf(rgbterm.FgString(fmt.Sprintf("Error occurred: %s", errMsg), 255, 0, 0))
	}
	return nil
}

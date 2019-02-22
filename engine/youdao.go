// Package engine Provides function to translate.
package engine

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/archervanderwaal/jadetrans/config"
	"github.com/archervanderwaal/jadetrans/utils"
	"github.com/aybabtme/rgbterm"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	api          = "http://openapi.youdao.com/api"
	appKeyLen    = 16
	appSecretLen = 32
)

// YoudaoEngine Represents a type that uses youdao translation.
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

// NewYoudaoEngine Returns an instance with translation functionality.
func NewYoudaoEngine(query, from, to, voice string, conf *config.Config) (*YoudaoEngine, error) {
	if len(query) < 1 {
		return nil, fmt.Errorf(fmt.Sprintf("Please enter the text to be translated"))
	}
	if len(strings.TrimSpace(conf.Youdao.AppKey)) != appKeyLen ||
					len(strings.TrimSpace(conf.Youdao.AppSecret)) != appSecretLen {
		return nil, fmt.Errorf(fmt.Sprintf(
			"Empty or incorrectly youdao appKey and appSecret: %s", conf.Youdao.AppKey))
	}
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
	return e, nil
}

// Query Returns formatted translate results.
func (e *YoudaoEngine) Query() string {
	spinnerID := utils.NewDefaultSpinnerAndStart("Querying...")
	resp, err := http.Get(e.requestURL())
	if err != nil {
		// err
		log.Println(rgbterm.BgString(err.Error(), 255, 0, 0))
		os.Exit(1)
	}
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	result := e.generateResult(res)
	formatRes := result.format()
	utils.StopSpinner(spinnerID)
	return formatRes
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

	// trans format
	if len(res.Basic.Explains) == 0 {
		for _, tran := range res.Translation {
			content += rgbterm.FgString(fmt.Sprintf("    %s\n", tran), 0, 255, 0)
		}
	}

	// format explains
	if len(res.Basic.Explains) != 0 {
		for _, exp := range res.Basic.Explains {
			content += rgbterm.FgString(fmt.Sprintf("    %s", exp), 0, 255, 0)
			content += "\n"
		}
		content += "\n"
	}

	if len(res.Web) != 0 {
		// web explains format
		for inx, w := range res.Web {
			content += rgbterm.FgString(fmt.Sprintf("    %d. %s\n", inx+1, w.Key), 0, 255, 0)
			content += rgbterm.FgString(fmt.Sprintf("       %s\n", strings.Join(w.Value, ", ")), 255, 0, 255)
		}
	}
	content += line
	return content
}

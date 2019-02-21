// Copyright 2019 Archer VanderWaal. All rights reserved.
// license that can be found in the LICENSE file.
package engine

const (
	URL 	=    "http://openapi.youdao.com/openapi"
)

type Result struct {
	ErrorCode   		string   `json:"errorCode"`
	Query       		string   `json:"query"`
	SpeakUrl			string	 `json:"speakUrl"`
	TranslateSpeakUrl	string	 `json:"tSpeakUrl"`
	Translation 		[]string `json:"translation"`
	Basic       		basic    `json:"basic"`
	Web         		[]web    `json:"web"`
}

type basic struct {
	Phonetic   			string   `json:"phonetic"`
	UkPhonetic 			string   `json:"uk-phonetic"`
	UsPhonetic 			string   `json:"us-phonetic"`
	Explains   			[]string `json:"explains"`
}

type web struct {
	Key   				string   `json:"key"`
	Value 				[]string `json:"value"`
}
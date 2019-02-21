package engine

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
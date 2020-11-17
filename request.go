package marusia

type Request struct {
	Meta struct {
		ClientId   string `json:"client_id"`
		Locale     string `json:"locale"`
		Timezone   string `json:"timezone"`
		Interfaces struct {
			Screen struct{} `json:"screen"`
		} `json:"interfaces"`
	} `json:"meta"`

	Request struct {
		Command           string `json:"command"`
		OriginalUtterance string `json:"original_utterance"`
		Type              string `json:"type"`
		Nlu               struct {
			Tokens   []string `json:"tokens"`
			Entities []string `json:"entities"`
		} `json:"nlu"`
	} `json:"request"`

	Session struct {
		SessionId string `json:"session_id"`
		UserId    string `json:"user_id"`
		SkillId   string `json:"skill_id"`
		New       bool   `json:"new"`
		MessageId int64  `json:"message_id"`
	} `json:"session"`

	Version string `json:"version"`
}

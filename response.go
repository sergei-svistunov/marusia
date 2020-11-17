package marusia

type FullResponse struct {
	Response Response        `json:"response"`
	Session  ResponseSession `json:"session"`
	Version  string          `json:"version"`
}

type Response struct {
	Text       string   `json:"text"`
	Tts        string   `json:"tts,omitempty"`
	Buttons    []Button `json:"buttons,omitempty"`
	Card       *Card    `json:"card,omitempty"`
	EndSession bool     `json:"end_session"`
}

type Button struct {
	Title   string      `json:"title"`
	Payload interface{} `json:"payload,omitempty"`
	Url     string      `json:"url,omitempty"`
}

type Card struct {
	Type        CardType   `json:"type"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	ImageId     int        `json:"image_id,omitempty"`
	Items       []CardItem `json:"items,omitempty"`
}

type CardType string

type CardItem struct {
	ImageId int `json:"image_id"`
}

type ResponseSession struct {
	SessionId string `json:"session_id"`
	MessageId int64  `json:"message_id"`
	UserId    string `json:"user_id"`
}

const (
	CardTypeBigImage  CardType = "BigImage"
	CardTypeItemsList CardType = "ItemsList"
)

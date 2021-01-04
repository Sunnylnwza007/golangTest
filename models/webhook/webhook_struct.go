package models

type Test struct {
	Events []Event `json:"events"`
	Destination string `json:"destination"`
}
type Event struct{
	Type string `json:"Type"`
	ReplyToken string `json:"replyToken"`
	Source Source `json:"source"`
	Timestamp int64 `json:"timestamp"`
	Mode string `json:"mode"`
	Message Message `json:"message"`

}

type Source struct{
	UserId string `json:"userId"`
	Type string `json:"type"`
	Address string `json:"address"`
}

type Message struct{
	Type string `json:"type"`
	Id string `json:"id"`
	Address string `json:"address"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
}
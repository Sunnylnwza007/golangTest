package models

type Test struct {
	Events []Event
	Destination string
}
type Event struct{
	Type string
	ReplyToken string
	Source Source
	Timestamp int64
	Mode string
	Message Message

}

type Source struct{
	UserId string
	Type string
	Address string
	Latitude float32
	Longitude float32
}


type Message struct{
	Type string
	Id string
	
}
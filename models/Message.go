package models

type Message struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewMessage(status int, msg string) Message {
	message := Message{}
	message.Status = status
	message.Message = msg
	return message
}

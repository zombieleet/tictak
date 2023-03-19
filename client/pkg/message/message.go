package message

import (
	"strings"
)

type Message struct{}

func InitMessage() *Message {
	return &Message{}
}

func (msg *Message) ProcessMessage(rawMessage string) (string, string) {
	splittedRawMessage := strings.Split(rawMessage, "<-")
	rawCommand := splittedRawMessage[0]
	return rawCommand[2:len(rawCommand)], strings.ReplaceAll(splittedRawMessage[1], "_", "\n")[0 : len(splittedRawMessage[1])-2]
}

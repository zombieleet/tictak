package message

import (
	"encoding/json"
	"github.com/zombieleet/tictak/client/pkg/message"
	"github.com/zombieleet/tictak/server/pkg/logger"
)

type messageContext struct {
	Type         string `json:"type"`
	ContextValue string `json:"value"`
	// will be use for broadcast
	Dst string `json:"dst,omitempty"`
	// will be use for broadcast
	Src string `json:"src,omitempty"`
}

type Message struct {
	Broadcast *broadcast
	Unicast   *unicast
	m         *message.Message
}

type MessageOptions struct {
	Logger *logger.Logger
}

func InitMessage(messageOptions MessageOptions) *Message {
	return &Message{
		Broadcast: &broadcast{messageOptions.Logger},
		Unicast:   &unicast{messageOptions.Logger},
		m:         message.InitMessage(),
	}
}

func (msg *Message) ProcessMessage(data string) (string, string) {
	return msg.m.ProcessMessage(data)
}

func encodeContext(mContext messageContext) string {
	jsonData, _ := json.Marshal(mContext)
	return string(jsonData)
}

package message

import (
	"encoding/json"
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
}

type MessageOptions struct {
	Logger *logger.Logger
}

func InitMessage(messageOptions MessageOptions) *Message {
	return &Message{
		Broadcast: &broadcast{messageOptions.Logger},
		Unicast:   &unicast{messageOptions.Logger},
	}
}

func encodeContext(mContext messageContext) string {
	jsonData, _ := json.Marshal(mContext)
	return string(jsonData)
}

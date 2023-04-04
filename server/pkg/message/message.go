package message

import (
	"encoding/json"
	"github.com/zombieleet/tictak/client/pkg/message"
	"github.com/zombieleet/tictak/server/pkg/logger"
	"github.com/zombieleet/tictak/server/pkg/players"
)

type messageContext struct {
	Type         string `json:"type"`
	ContextValue string `json:"value"`
	// it will be use for sending unicast messages
	Dst string `json:"dst"`
	// sender address
	Src string `json:"src"`
	// it will be use for sending broadcast messages
	DstMultiple string `json:dst_multiple`
}

type Message struct {
	Broadcast *broadcast
	Unicast   *unicast
	m         *message.Message
}

type MessageOptions struct {
	Logger                       *logger.Logger
	ConnectedPlayersForBroadcast *players.PlayerConnectedToGame
}

func InitMessage(messageOptions MessageOptions) *Message {
	return &Message{
		Broadcast: &broadcast{messageOptions.Logger, messageOptions.ConnectedPlayersForBroadcast},
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

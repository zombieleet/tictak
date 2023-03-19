package handlers

import (
	"github.com/zombieleet/tictak/server/pkg/room"
)

type HandlersMap map[string]func(...interface{})

type Handler struct {
	handlers HandlersMap
}

type HandlerOption struct {
	Room room.Room
}

func InitHandlers(handlerOption HandlerOption) *Handler {

	handler := &Handler{
		handlers: make(HandlersMap),
	}

	handler.handlers["CMD_ENTER_ROOM"] = handlerOption.Room.EnterRoom
	return handler
}

func (handler *Handler) HandleCommand(command, payload, address string) {
	handlerFunc, ok := handler.handlers[command]

	// TODO: log error and return
	// send client an error
	if !ok {
		return
	}

	handlerFunc(payload, address)
}

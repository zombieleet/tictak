package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/zombieleet/tictak/server/pkg/room"
	"net"
)

type HandlersMap map[string]func(net.Conn, chan string, context.CancelCauseFunc, ...interface{})

type Handler struct {
	handlers      HandlersMap
	cancelCtxFunc context.CancelCauseFunc
	commsChan     chan string
}

type HandlerOption struct {
	Room          room.Room
	CancelCtxFunc context.CancelCauseFunc
	CommsChan     chan string
}

func InitHandlers(handlerOption HandlerOption) *Handler {

	handler := &Handler{
		handlers:      make(HandlersMap),
		cancelCtxFunc: handlerOption.CancelCtxFunc,
		commsChan:     handlerOption.CommsChan,
	}

	handler.handlers["CMD_ENTER_ROOM"] = handlerOption.Room.EnterRoom
	return handler
}

func (handler *Handler) HandleCommand(command, payload, address string, conn net.Conn) {
	handlerFunc, ok := handler.handlers[command]

	if !ok {
		handler.cancelCtxFunc(errors.Join(E_HANDLER_NO_EXIST, fmt.Errorf("for command %s", command)))
		return
	}

	handlerFunc(conn, handler.commsChan, handler.cancelCtxFunc, payload, address)
}

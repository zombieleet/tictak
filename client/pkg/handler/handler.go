package handler

import (
	"context"
	"errors"
	"github.com/zombieleet/tictak/client/pkg/gclienterror"
	"github.com/zombieleet/tictak/client/pkg/ui"
)

// Handler struct holds private and public fields needed for the handler package
// the handler package delegates operations to the UI
type Handler struct {
	// relationship between server commands and handler functions
	handlers map[string]func(chan string, context.Context, string)
	ui       *ui.UI

	// context for closing the connection when an error that cannot
	// be recovered from occurs
	ctx            context.Context
	cancelCtxCause context.CancelCauseFunc

	// channel for communicating between the connection package and the ui package
	// handler package passes the channel to to methods in the UI package
	commsChan chan string
}

// InitHandlers maps the server command with ui handlers
func InitHandlers(ctx context.Context, cancelCtxCause context.CancelCauseFunc, commsChan chan string) *Handler {
	handler := &Handler{
		handlers:       make(map[string]func(chan string, context.Context, string)),
		ui:             ui.CreateUI(cancelCtxCause),
		ctx:            ctx,
		commsChan:      commsChan,
		cancelCtxCause: cancelCtxCause,
	}

	handler.handlers["CMD_SEND_ROOMS"] = handler.ui.Room.CreateRoomListUI
	handler.handlers["CMD_UPDATE_ROOMS"] = handler.ui.Room.UpdateRoomListUI

	return handler
}

// CreateUI dynamically invokes UI updates/creation/deletion
// it works with the `Handler.Handlers` map
func (handler *Handler) HandleUICommand(command string, payload string) {

	handlerFunc, ok := handler.handlers[command]

	if !ok {
		handler.cancelCtxCause(errors.Join(gclienterror.HANDLER_FUNC_ERROR, errors.New("cannot find handler command: "+command)))
		return
	}

	handlerFunc(handler.commsChan, handler.ctx, payload)

	if err := handler.ui.App.SetRoot(handler.ui.MainLayout, true).EnableMouse(true).Run(); err != nil {
		handler.cancelCtxCause(errors.Join(gclienterror.UI_STARTUP_ERROR, err))
		return
	}

}

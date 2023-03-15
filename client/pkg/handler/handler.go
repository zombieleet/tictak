package handler

import (
	"context"
	"errors"
	"github.com/zombieleet/tictak/client/pkg/gclienterror"
	"github.com/zombieleet/tictak/client/pkg/ui"
)

type Handler struct {
	Handlers       map[string]func(context.Context, string)
	ui             *ui.UI
	ctx            context.Context
	cancelCtxCause context.CancelCauseFunc
}

func InitHandlers(ctx context.Context, cancelCtxCause context.CancelCauseFunc, commsChan chan any) *Handler {
	handler := &Handler{
		Handlers:       make(map[string]func(context.Context, string)),
		ui:             ui.CreateUI(cancelCtxCause),
		ctx:            ctx,
		cancelCtxCause: cancelCtxCause,
	}

	handler.Handlers["CMD_SEND_ROOMS"] = handler.ui.Room.CreateRoomListUI
	handler.Handlers["CMD_UPDATE_ROOMS"] = handler.ui.Room.UpdateRoomListUI

	return handler
}

// CreateUI
// This responsible for dynamically invoking the UI functionality
func (handler *Handler) HandleUICommand(command string, payload string) {

	handlerFunc, ok := handler.Handlers[command]

	if !ok {
		handler.cancelCtxCause(errors.Join(gclienterror.HANDLER_FUNC_ERROR, errors.New("cannot find handler command: "+command)))
		return
	}

	handlerFunc(handler.ctx, payload)

	if err := handler.ui.App.SetRoot(handler.ui.Layout, true).EnableMouse(true).Run(); err != nil {
		handler.cancelCtxCause(errors.Join(gclienterror.UI_STARTUP_ERROR, err))
		return
	}

}

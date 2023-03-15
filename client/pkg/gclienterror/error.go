package gclienterror

import "errors"

var (
	HANDLER_FUNC_ERROR = errors.New("HANDLER_FUNC_ERROR")
	UI_STARTUP_ERROR   = errors.New("UI_STARTUP_ERROR")
)

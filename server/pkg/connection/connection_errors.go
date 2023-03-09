package connection

import (
	"errors"
)

var (
	E_NOT_RESOLVED = errors.New("Cannot resolve address")
	E_NOT_ACCEPT   = errors.New("Cannot accept connection")
)

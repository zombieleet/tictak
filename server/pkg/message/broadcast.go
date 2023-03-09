package message

import (
	"github.com/zombieleet/tictak/server/pkg/logger"
)

type broadcast struct {
	logger *logger.Logger
}

func (bcast *broadcast) NO_OP() {}

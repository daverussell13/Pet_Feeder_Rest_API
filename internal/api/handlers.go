package api

import (
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/realtime"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/schedule"
)

type V1Handlers struct {
	realtime realtime.Handler
	schedule schedule.Handler
}

type Handlers struct {
	V1 *V1Handlers
}

func NewHandler(v1 *V1Handlers) *Handlers {
	return &Handlers{
		v1,
	}
}

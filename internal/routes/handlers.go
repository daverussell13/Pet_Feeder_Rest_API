package routes

import (
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/realtime"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/schedule"
)

type APIV1Handlers struct {
	realtime realtime.Handler
	schedule schedule.Handler
}

type Handlers struct {
	V1 *APIV1Handlers
}

func NewHandler(v1 *APIV1Handlers) *Handlers {
	return &Handlers{
		v1,
	}
}

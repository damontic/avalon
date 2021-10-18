package metrics

import (
	"runtime"

	"github.com/damontic/avalon/internal/domain/room"
	"github.com/damontic/avalon/internal/domain/state"
)

type MetricsSummary struct {
	Rooms          map[int]room.Room `json:"rooms"`
	MaxNumberRooms int               `json:"maxNumberRooms"`
	GoVersion      string            `json:"goVersion"`
}

func GetMetricsSummary(state *state.State) *MetricsSummary {
	rooms := state.Rooms
	maxNumberRooms := state.MaxNumberRooms
	return &MetricsSummary{
		Rooms:          rooms,
		MaxNumberRooms: maxNumberRooms,
		GoVersion:      runtime.Version(),
	}
}

package metrics

import (
	"fmt"
	"runtime"
	"time"

	"github.com/damontic/avalon/internal/domain/room"
	"github.com/damontic/avalon/internal/domain/state"
)

type MetricsSummary struct {
	Rooms            map[int]room.Room `json:"rooms"`
	MaxNumberRooms   int               `json:"maxNumberRooms"`
	GoVersion        string            `json:"goVersion"`
	MemoryAllocation string            `json:"memoryAllocation"`
	TotalAllocation  string            `json:"totalAllocation"`
	SysMemory        string            `json:"sysMemory"`
	NumGC            uint32            `json:"numGC"`
	Date             string            `json:"date"`
	AvalonVersion    string            `json:"avalonVersion"`
}

func GetMetricsSummary(state *state.State) *MetricsSummary {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	rooms := state.Rooms
	maxNumberRooms := state.MaxNumberRooms
	version := state.Version

	return &MetricsSummary{
		Date:             time.Now().UTC().Format(time.RFC3339),
		GoVersion:        runtime.Version(),
		AvalonVersion:    version,
		Rooms:            rooms,
		MaxNumberRooms:   maxNumberRooms,
		MemoryAllocation: bToMb(m.Alloc),
		TotalAllocation:  bToMb(m.TotalAlloc),
		SysMemory:        bToMb(m.Sys),
		NumGC:            m.NumGC,
	}
}

func bToMb(b uint64) string {
	var megas uint64 = b / 1024 / 1024
	return fmt.Sprintf("%d MiB", megas)
}

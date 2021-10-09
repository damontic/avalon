package tests

import (
	"testing"

	avalon "github.com/damontic/avalon/internal/server"
)

func TestNewAvalonServer(t *testing.T) {
	server := avalon.NewAvalonServer(10, 80, false, "", "", "", false)
	if len(server.Rooms) != 0 {
		t.Fatalf("Expected Number of Rooms is 0, got: %d", len(server.Rooms))
	}
}

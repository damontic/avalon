package main

import (
	avalon "github.com/damontic/avalon/internal"
)

func main() {
	server := avalon.NewAvalonServer()
	server.Run()
}

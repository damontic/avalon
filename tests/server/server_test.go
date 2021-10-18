package tests

import (
	"fmt"

	avalon "github.com/damontic/avalon/internal/server"
)

func ExampleAvalonServer() {
	server := avalon.NewAvalonServer("0.0.0", 10, 80, false, "", "", "", false, 4, false)
	fmt.Printf("%s", server)
	// Output:
	// {"port":80,"isSsl":false,"isHttpToHttpsEnabled":false,"sslCert":"","sslKey":"","domain":"","state":{"version":"0.0.0","rooms":{},"maxNumberRooms":10}}
}

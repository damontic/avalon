package tests

import (
	"fmt"

	"github.com/damontic/avalon/internal/server/jsend"
)

func ExampleJsendResponse() {
	resp := jsend.JsendResponse{
		Success: true,
		Data:    1,
	}
	fmt.Printf("%s\n", resp)

	resp = jsend.JsendResponse{
		Success: false,
		Error:   "Bad thing happened",
	}
	fmt.Printf("%s", resp)
	// Output:
	// {"success":true,"data":1}
	// {"success":false,"error":"Bad thing happened"}
}

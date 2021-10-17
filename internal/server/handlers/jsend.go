package handlers

import (
	"encoding/json"
	"log"
)

type JsendResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (jsr JsendResponse) String() string {
	res, err := json.Marshal(jsr)
	if err != nil {
		log.Fatalln("Error Marshalling AvalonServer")
	}
	return string(res[:])
}

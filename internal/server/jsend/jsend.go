package jsend

import (
	"encoding/json"
	"log"
)

type JsendResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewJsendResponseSuccess() JsendResponse {
	return JsendResponse{
		Success: true,
	}
}

func NewJsendResponseSuccessData(data interface{}) JsendResponse {
	return JsendResponse{
		Success: true,
		Data:    data,
	}
}

func NewJsendResponseFailure(message string) JsendResponse {
	return JsendResponse{
		Success: false,
		Error:   message,
	}
}

func (jsr JsendResponse) String() string {
	res, err := json.Marshal(jsr)
	if err != nil {
		log.Fatalln("Error Marshalling AvalonServer")
	}
	return string(res[:])
}

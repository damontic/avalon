package handlers

type JsendResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Code    int    `json:"code,omitempty"`
}

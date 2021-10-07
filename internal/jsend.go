package avalon

type JsendResponse struct {
	Success string `json:"Success"`
	Failure string `json:"Failure"`
	Error   string `json:"Error"`
	Code    int    `json:"Code"`
}

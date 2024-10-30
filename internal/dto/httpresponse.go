package dto

type HttpResponse struct {
	Result  any    `json:"result"`
	Error   string `json:"error,omitempty"`
	Success bool   `json:"success"`
}

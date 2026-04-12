package config

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
	Detail     string `json:"detail"`
}

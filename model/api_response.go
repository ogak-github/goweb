package model

type ApiResponse struct {
	Code          int    `json:"code"`
	StatusMessage string `json:"status_message"`
	Data          any    `json:"data"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	ExpiredIn string `json:"expired_in"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}

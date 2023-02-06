package dto

type LoginRequestForm struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token,omitempty"`
}

package entities

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token       string `json:"token"`
	RefresToken string `json:"refres_token,omitempty"`
	User        *User  `json:"user,omitempty"`
}

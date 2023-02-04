package models

/* LoginResponse login response structure */
type LoginResponse struct {
	JwtToken string `json:"token,omitempty"`
}

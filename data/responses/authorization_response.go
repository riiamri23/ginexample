package responses

type AuthorizationResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

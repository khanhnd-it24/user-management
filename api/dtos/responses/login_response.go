package responses

type LoginResponse struct {
	Data        UserResponse `json:"data"`
	AccessToken string       `json:"accessToken"`
}

package requests

type UserRequest struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

func Validate(userRequest UserRequest) bool {
	if userRequest.Username == nil || len(*userRequest.Username) < 6 {
		return false
	}

	if userRequest.Password == nil || len(*userRequest.Password) < 8 {
		return false
	}

	return true
}

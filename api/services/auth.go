package services

import (
	"github.com/gin-gonic/gin"
	"user-management/adapter/repositories"
	"user-management/api/dtos/requests"
	"user-management/api/dtos/responses"
	"user-management/core/entities"
	"user-management/core/enums"
	"user-management/utils"
)

type AuthService interface {
	Register(data requests.UserRequest) (responses.UserResponse, error)
	Login(data requests.UserRequest) (responses.LoginResponse, error)
}

type authService struct {
	userRepository repositories.UserRepository
}

func New(ur repositories.UserRepository) AuthService {
	return &authService{
		userRepository: ur,
	}
}

func (a authService) Register(data requests.UserRequest) (responses.UserResponse, error) {
	var existedUser entities.User
	var err error
	existedUser, err = a.userRepository.FindUserByUsername(
		*data.Username)

	if err != nil || existedUser.Id != 0 {
		return responses.UserResponse{}, gin.Error{Meta: "user already existed", Type: 400}
	}

	password, _ := utils.HashPassword([]byte(*data.Password))

	user := entities.User{
		Username: *data.Username,
		Password: password,
		Role:     enums.SuperAdmin.String(),
	}

	user, _ = a.userRepository.Save(user)

	userResponse := responses.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Role:     user.Role,
	}

	return userResponse, nil
}

func (a authService) Login(data requests.UserRequest) (responses.LoginResponse, error) {
	var existedUser entities.User
	var err error
	existedUser, err = a.userRepository.FindUserByUsername(*data.Username)

	if err != nil || existedUser.Id == 0 {
		return responses.LoginResponse{}, gin.Error{Meta: "User not found", Type: 404}
	}

	err = utils.CompareHashAndPassword(existedUser.Password, []byte(*data.Password))

	if err != nil {
		return responses.LoginResponse{}, gin.Error{Meta: "Password incorrect", Type: 400}
	}

	user := responses.UserResponse{
		Id:       existedUser.Id,
		Username: existedUser.Username,
		Role:     existedUser.Role,
	}

	var token *string
	token, err, _ = utils.GenerateTokenAndHandleError(existedUser, 30)

	userRes := responses.LoginResponse{Data: user, AccessToken: *token}
	return userRes, nil
}

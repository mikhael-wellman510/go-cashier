package usecases

import (
	"errors"
	"log"
	"mikhael-project-go/internal/adapters/repositories"
	"mikhael-project-go/internal/entities"
	"mikhael-project-go/internal/utils"
)

type (
	AuthService interface {
		RegisterUser(userReq *entities.UserRequest) (*entities.UserResponse, error)
		Login(loginReq *entities.LoginRequest) (*entities.LoginResponse, error)
	}

	authService struct {
		AuthRepository repositories.UserRepository
	}
)

func NewAuthService(userRepository repositories.UserRepository) AuthService {

	return &authService{
		AuthRepository: userRepository,
	}
}

func (as *authService) Login(loginReq *entities.LoginRequest) (*entities.LoginResponse, error) {
	// Todo -> Check email
	user, err := as.AuthRepository.FindByEmail(loginReq.Email)

	if err != nil {
		return nil, errors.New("EMAIL OR PASSWORD INVALID 1")
	}

	// chack pass hash

	check := utils.CheckPasswordHash(user.Password, loginReq.Password)

	if !check {
		return nil, errors.New("EMAIL OR PASSWORD INVALID 2")
	}

	// Generate token

	return nil, nil
}
func (as *authService) RegisterUser(userReq *entities.UserRequest) (*entities.UserResponse, error) {

	// hash password
	passwordHash, err := utils.HashPassword(userReq.Password)

	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Username:    userReq.Username,
		Email:       userReq.Email,
		Password:    passwordHash,
		PhoneNumber: userReq.PhoneNumber,
		Address:     userReq.Address,
		IsActive:    true,
	}

	if err := as.AuthRepository.Create(user); err != nil {
		return nil, err
	}

	log.Println("Hasil user : ", user)
	return &entities.UserResponse{
		Id:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		IsActive:    user.IsActive,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}

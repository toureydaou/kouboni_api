package services

import (
	u "auth-service/internal/models"
	"auth-service/internal/repository"
	"auth-service/internal/utils"
	"context"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterUser(ctx context.Context, userRegister u.UserRegister) (response *u.UserResponse, error error) {

	hashedPassword := utils.HashPassword(userRegister.Password)

	user := u.User{
		Name:        userRegister.Name,
		Surname:     userRegister.Surname,
		Gender:      userRegister.Gender,
		Email:       userRegister.Email,
		Password:    hashedPassword,
		PhoneNumber: userRegister.PhoneNumber,
	}

	existing, _ := s.repo.FindUserByEmail(ctx, user.Email)

	if existing != nil {
		return nil, utils.ErrorEmailExists
	}

	existing, _ = s.repo.FindUserByPhoneNumber(ctx, user.PhoneNumber)

	if existing != nil {
		return nil, utils.ErrorPhoneNumberExists
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return utils.ToUserResponse(user), nil
}

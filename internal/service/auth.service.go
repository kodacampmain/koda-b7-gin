package service

import (
	"context"

	"github.com/kodacampmain/koda-b7-gin/internal/dto"
	"github.com/kodacampmain/koda-b7-gin/internal/repository"
	"github.com/kodacampmain/koda-b7-gin/pkg"
)

type AuthService struct {
	authRepo *repository.AuthRepository
}

func NewAuthService(authRepo *repository.AuthRepository) *AuthService {
	return &AuthService{
		authRepo: authRepo,
	}
}

func (a *AuthService) RegisterUser(ctx context.Context, user dto.NewUser) (dto.User, error) {
	var hc pkg.HashConfig
	hc.UseRecommended()
	hashedPwd := hc.GenHash(user.Password)
	newUser, err := a.authRepo.AddNewUser(ctx, user.Username, hashedPwd)
	if err != nil {
		return dto.User{}, err
	}
	return dto.User{
		Id:        newUser.Id,
		Username:  newUser.Username,
		Password:  newUser.Password,
		CreatedAt: newUser.CreatedAt,
	}, nil
}

func (a *AuthService) LoginUser(ctx context.Context, user dto.NewUser) (string, error) {
	login, err := a.authRepo.GetUserByUsername(ctx, user.Username)
	if err != nil {
		return "", err
	}
	var hc pkg.HashConfig

	if err := hc.Compare(user.Password, login.Password); err != nil {
		return "", err
	}
	claims := pkg.NewClaims(login.Id, user.Username)
	token, err := claims.GenJWT()
	if err != nil {
		return "", err
	}
	return token, nil
}

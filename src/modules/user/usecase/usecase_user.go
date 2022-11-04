package usecase

import (
	"context"
	"graphql_go/src/modules/user/domain"
	"graphql_go/src/modules/user/repository"
)

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return userUsecase{
		userRepo: userRepo,
	}
}

func (u userUsecase) Login(ctx context.Context, data domain.LoginPayload) (domain.LoginResponse, error) {
	return domain.LoginResponse{}, nil
}

func (u userUsecase) Register(ctx context.Context, data domain.RegisterPayload) error {
	return nil
}

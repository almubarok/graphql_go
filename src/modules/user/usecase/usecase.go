package usecase

import (
	"context"
	"graphql_go/src/modules/user/domain"
)

type UserUsecase interface {
	Login(ctx context.Context, data domain.LoginPayload) (domain.LoginResponse, error)
	Register(ctx context.Context, data domain.RegisterPayload) error
}

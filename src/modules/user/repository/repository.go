package repository

import (
	"context"

	"graphql_go/src/modules/user/domain"
	"graphql_go/src/modules/user/model"
)

type UserRepository interface {
	FindOne(ctx context.Context, filter domain.UserFilter) (model.User, error)
	Save(ctx context.Context, data *model.User) error
}

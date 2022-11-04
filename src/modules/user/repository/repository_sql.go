package repository

import (
	"context"
	"graphql_go/src/modules/user/domain"
	"graphql_go/src/modules/user/model"

	"gorm.io/gorm"
)

type userRepoSQL struct {
	db *gorm.DB
}

func NewUserRepoSQL(db *gorm.DB) UserRepository {
	return &userRepoSQL{db: db}
}

func (r userRepoSQL) FindOne(ctx context.Context, filter domain.UserFilter) (model.User, error) {

	return model.User{}, nil
}

func (r userRepoSQL) Save(ctx context.Context, data *model.User) error {
	return nil
}

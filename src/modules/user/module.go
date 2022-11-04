package user

import (
	"graphql_go/config"
	"graphql_go/src/modules/user/delivery/resthandler"
	"graphql_go/src/modules/user/repository"
	"graphql_go/src/modules/user/usecase"
	"graphql_go/src/shared"
)

type Module struct {
	restHandler    shared.RESTHandler
	graphqlHandler shared.GraphQLHandler
}

func NewModule(cfg *config.Config) *Module {
	var mod Module
	repo := repository.NewUserRepoSQL(cfg.DBConn)
	uc := usecase.NewUserUsecase(repo)

	mod.restHandler = resthandler.NewRestHandler(uc)

	return &mod
}

func (m Module) RESTHandler() shared.RESTHandler {
	return m.restHandler
}
func (m Module) GraphQLHandler() shared.GraphQLHandler {
	return m.graphqlHandler
}

package modules

import (
	"graphql_go/config"
	"graphql_go/src/modules/user"
	"graphql_go/src/shared"
)

type ModuleFactory interface {
	RESTHandler() shared.RESTHandler
	GraphQLHandler() shared.GraphQLHandler
}

type ServiceFactory interface {
	GetConfig() *config.Config
	GetModules() []ModuleFactory
}

type Service struct {
	config  *config.Config
	modules []ModuleFactory
}

func NewService(cfg *config.Config) ServiceFactory {
	modules := []ModuleFactory{
		user.NewModule(cfg),
	}

	return Service{
		config:  cfg,
		modules: modules,
	}
}

func (s Service) GetConfig() *config.Config {
	return s.config
}
func (s Service) GetModules() []ModuleFactory {
	return s.modules
}

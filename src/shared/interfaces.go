package shared

import (
	echo "github.com/labstack/echo/v4"
)

type RESTHandler interface {
	Mount(group *echo.Group)
}

type GraphQLHandler interface {
	Query() interface{}
	Mutation() interface{}
	Subscription() interface{}
	RegisterMiddleware() interface{}
}

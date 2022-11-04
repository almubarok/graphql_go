package resthandler

import (
	"graphql_go/src/modules/user/usecase"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type RestHandler struct {
	userUsecase usecase.UserUsecase
}

func NewRestHandler(userUsecase usecase.UserUsecase) *RestHandler {
	return &RestHandler{userUsecase: userUsecase}
}

func (h RestHandler) Mount(root *echo.Group) {
	v1Root := root.Group("/v1")

	user := v1Root.Group("/user")
	user.GET("/ping", h.ping)
}

func (h RestHandler) ping(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}

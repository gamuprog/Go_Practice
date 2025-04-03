package router

import (
	"Go_Practice/api/handler"
	"Go_Practice/usecase/interactor"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer(
	userUC interactor.IUserUseCase,
) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	userHandler := handler.NewUserHandler(userUC)


	api := e.Group("/api")

	api.GET("/:userId", userHandler.FindById)

	return e
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muharanii/auth-rest-api-go/internal/api"
	"github.com/muharanii/auth-rest-api-go/internal/component"
	"github.com/muharanii/auth-rest-api-go/internal/config"
	"github.com/muharanii/auth-rest-api-go/internal/middleware"
	"github.com/muharanii/auth-rest-api-go/internal/repository"
	"github.com/muharanii/auth-rest-api-go/internal/service"
)


func main() {
	cnf := config.Get()
	dbConnection := component.GetDatabaseConnection(cnf)
	cacheConnection := component.GetCacheConnection()

	userRepository := repository.NewUser(dbConnection)

	userService := service.NewUser(userRepository, cacheConnection)

	authMid := middleware.Authenticate(userService)

	app := fiber.New()
	api.NewAuth(app, userService, authMid)

	_ = app.Listen(cnf.Server.Host+":"+cnf.Server.Port)
}

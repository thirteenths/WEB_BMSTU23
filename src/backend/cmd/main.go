package main

// @title Go + Gin Todo API
// @version 1.0
// @description This is a sample server todo server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8000
// @BasePath /
// @query.collection.format multi
import (
	"github.com/sirupsen/logrus"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/handler"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/service"
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/server"
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/storage/postgres"
)

func main() {

	storage, err := postgres.NewPostgresStorage("standup", "password", "standup")
	if err != nil {
		logrus.Fatalf("error %s", err.Error())
	}

	serv := service.NewService(storage)
	handlers := handler.NewHandler(serv)

	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error %s", err.Error())
	}

}

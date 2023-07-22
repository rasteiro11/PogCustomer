package main

import (
	middlewares "flashcards/middleware"
	"flashcards/models"
	usersHttp "flashcards/src/user/delivery/http"
	usersRepo "flashcards/src/user/repository"
	usersCase "flashcards/src/user/usecase"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rasteiro11/PogCore/pkg/database"
	"github.com/rasteiro11/PogCore/pkg/server"
	"log"
)

func main() {
	database, err := database.NewDatabase(database.GetMysqlEngineBuilder)
	if err != nil {
		log.Fatalf("[main] database.NewDatabase() retunrned error: %+v\n", err)
	}

	if err := database.Migrate(models.GetEntities()...); err != nil {
		log.Fatalf("[main] database.Migrate() retunrned error: %+v\n", err)
	}

	server := server.NewServer(server.WithPrefix("/customer"))
	server.Use("/user", middlewares.ValidateUserMiddleware())
	server.Use("/*", cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	db := database.Conn()
	usersRepo := usersRepo.NewRepository(db)

	usersUsecase := usersCase.NewUsecase(
		usersCase.WithRepository(usersRepo),
	)

	usersHttp.NewHandler(server, usersHttp.WithUsecase(usersUsecase))

	server.PrintRouter()

	if err := server.Start(":6969"); err != nil {
		log.Fatalf("[main] server.NewServer() retrurned error: %+v\n", err)
	}
}

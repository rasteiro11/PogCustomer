package main

import (
	"log"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rasteiro11/PogCore/pkg/database"
	"github.com/rasteiro11/PogCore/pkg/server"
	"github.com/rasteiro11/PogCustomer/entities"
	middlewares "github.com/rasteiro11/PogCustomer/middleware"
	usersHttp "github.com/rasteiro11/PogCustomer/src/user/delivery/http"
	usersRepo "github.com/rasteiro11/PogCustomer/src/user/repository"
	usersCase "github.com/rasteiro11/PogCustomer/src/user/usecase"
)

func main() {
	database, err := database.NewDatabase(database.GetMysqlEngineBuilder)
	if err != nil {
		log.Fatalf("[main] database.NewDatabase() retunrned error: %+v\n", err)
	}

	if err := database.Migrate(entities.GetEntities()...); err != nil {
		log.Fatalf("[main] database.Migrate() returned error: %+v\n", err)
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
		log.Fatalf("[main] server.NewServer() returned error: %+v\n", err)
	}
}

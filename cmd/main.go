package main

import (
	"context"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rasteiro11/PogCore/pkg/config"
	"github.com/rasteiro11/PogCore/pkg/database"
	"github.com/rasteiro11/PogCore/pkg/logger"
	"github.com/rasteiro11/PogCore/pkg/server"
	"github.com/rasteiro11/PogCore/pkg/transport/grpcserver"
	"github.com/rasteiro11/PogCustomer/entities"
	pbCrypto "github.com/rasteiro11/PogCustomer/gen/proto/go/crypto"
	pbCustomer "github.com/rasteiro11/PogCustomer/gen/proto/go/customer"
	middlewares "github.com/rasteiro11/PogCustomer/middleware"
	usersHttp "github.com/rasteiro11/PogCustomer/src/user/delivery/http"
	usersRepo "github.com/rasteiro11/PogCustomer/src/user/repository"
	usersSvc "github.com/rasteiro11/PogCustomer/src/user/service"
	usersCase "github.com/rasteiro11/PogCustomer/src/user/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	database, err := database.NewDatabase(database.GetMysqlEngineBuilder)
	if err != nil {
		logger.Of(ctx).Fatalf("[main] database.NewDatabase() retunrned error: %+v\n", err)
	}

	if err := database.Migrate(entities.GetEntities()...); err != nil {
		logger.Of(ctx).Fatalf("[main] database.Migrate() returned error: %+v\n", err)
	}

	server := server.NewServer(server.WithPrefix("/customer"))
	server.Use("/user", middlewares.ValidateUserMiddleware())
	server.Use("/*", cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	db := database.Conn()

	usersRepo := usersRepo.NewRepository(db)

	credentials := insecure.NewCredentials()
	cryptoConn, err := grpc.Dial(config.Instance().String("CRYPTO_SERVICE"),
		grpc.WithTransportCredentials(credentials))
	if err != nil {
		logger.Of(ctx).Fatalf(
			"[main] grpc.Dial returned error: err=%+v", err)
	}

	cryptoClient := pbCrypto.NewCryptoServiceClient(cryptoConn)

	usersUsecase := usersCase.NewUsecase(
		usersCase.WithCryptoClient(cryptoClient),
		usersCase.WithRepository(usersRepo),
	)

	customerSvc := usersSvc.NewService(usersSvc.WithUserUsecase(usersUsecase))

	go func() {
		server := grpcserver.NewServer(grpcserver.WithReflectionEnabled())

		server.Register(pbCustomer.CustomerService_ServiceDesc, customerSvc)

		if err := server.Run(); err != nil {
			logger.Of(ctx).Fatalf("[main] server.Run() returned error: %+v", err)
		}
	}()

	usersHttp.NewHandler(server, usersHttp.WithUsecase(usersUsecase))

	server.PrintRouter()

	if err := server.Start("192.168.0.14:6969"); err != nil {
		logger.Of(ctx).Fatalf("[main] server.NewServer() returned error: %+v\n", err)
	}
}

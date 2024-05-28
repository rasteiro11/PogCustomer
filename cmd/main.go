package main

import (
	"context"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rasteiro11/PogCore/pkg/database"
	"github.com/rasteiro11/PogCore/pkg/logger"
	"github.com/rasteiro11/PogCore/pkg/server"
	"github.com/rasteiro11/PogCore/pkg/transport/grpcserver"
	"github.com/rasteiro11/PogPaymentSheet/entities"
	defaultVrRepo "github.com/rasteiro11/PogPaymentSheet/src/defaultvr/repository"
	departmentRepo "github.com/rasteiro11/PogPaymentSheet/src/department/repository"
	employeeRepo "github.com/rasteiro11/PogPaymentSheet/src/employee/repository"
	employeeFrequencyRepo "github.com/rasteiro11/PogPaymentSheet/src/employeefrequency/repository"
	frequencyHttp "github.com/rasteiro11/PogPaymentSheet/src/frequency/delivery/http"
	frequencyRepo "github.com/rasteiro11/PogPaymentSheet/src/frequency/repository"
	rankRepo "github.com/rasteiro11/PogPaymentSheet/src/rank/repository"
	vrRepo "github.com/rasteiro11/PogPaymentSheet/src/vr/repository"

	// pbCrypto "github.com/rasteiro11/PogPaymentSheet/gen/proto/go/crypto"
	// pbCustomer "github.com/rasteiro11/PogPaymentSheet/gen/proto/go/customer"
	middlewares "github.com/rasteiro11/PogPaymentSheet/middleware"
	// roleRepo "github.com/rasteiro11/PogPaymentSheet/src/role/repository"
	departmentHttp "github.com/rasteiro11/PogPaymentSheet/src/department/delivery/http"
	employeeHttp "github.com/rasteiro11/PogPaymentSheet/src/employee/delivery/http"
	rankHttp "github.com/rasteiro11/PogPaymentSheet/src/rank/delivery/http"
	vrHttp "github.com/rasteiro11/PogPaymentSheet/src/vr/delivery/http"

	// usersRepo "github.com/rasteiro11/PogPaymentSheet/src/user/repository"
	// usersSvc "github.com/rasteiro11/PogPaymentSheet/src/user/service"
	employeeCase "github.com/rasteiro11/PogPaymentSheet/src/employee/usecase"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
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

	server := server.NewServer(server.WithPrefix("/payment/sheet"))
	server.Use("/user", middlewares.ValidateUserMiddleware())
	server.Use("/*", cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	db := database.Conn()

	defaultVrRepo := defaultVrRepo.NewRepository(db)
	vrRepo := vrRepo.NewRepository(db)
	departmentRepo := departmentRepo.NewRepository(db)
	rankRepo := rankRepo.NewRepository(db)
	employeeRepo := employeeRepo.NewRepository(db)
	frequencyRepo := frequencyRepo.NewRepository(db)
	employeeFrequencyRepo := employeeFrequencyRepo.NewRepository(db)

	// credentials := insecure.NewCredentials()
	// cryptoConn, err := grpc.Dial(config.Instance().String("CRYPTO_SERVICE"),
	// 	grpc.WithTransportCredentials(credentials))
	// if err != nil {
	// 	logger.Of(ctx).Fatalf(
	// 		"[main] grpc.Dial returned error: err=%+v", err)
	// }

	// cryptoClient := pbCrypto.NewCryptoServiceClient(cryptoConn)

	employeeUsecase := employeeCase.NewUsecase(
		employeeCase.WithDefaultVrRepository(defaultVrRepo),
		employeeCase.WithVrRepository(vrRepo),
		employeeCase.WithRankRepository(rankRepo),
		employeeCase.WithDefaultVrRepository(defaultVrRepo),
		employeeCase.WithDepartmentRepository(departmentRepo),
		employeeCase.WithRepository(employeeRepo),
		employeeCase.WithEmployeeFrequencyRepo(employeeFrequencyRepo),
		employeeCase.WithFrequencyRepo(frequencyRepo),
	)

	// _ = usersSvc.NewService(usersSvc.WithUserUsecase(usersUsecase))

	go func() {
		server := grpcserver.NewServer(grpcserver.WithReflectionEnabled())

		//	server.Register(pbCustomer.CustomerService_ServiceDesc, customerSvc)

		if err := server.Run(); err != nil {
			logger.Of(ctx).Fatalf("[main] server.Run() returned error: %+v", err)
		}
	}()

	employeeHttp.NewHandler(server, employeeHttp.WithUsecase(employeeUsecase))
	departmentHttp.NewHandler(server, departmentHttp.WithRepository(departmentRepo))
	rankHttp.NewHandler(server, rankHttp.WithRepository(rankRepo))
	vrHttp.NewHandler(server,
		vrHttp.WithRepository(vrRepo),
		vrHttp.WithEmployeeRepo(employeeRepo))
	frequencyHttp.NewHandler(server,
		frequencyHttp.WithEmployeeRepo(employeeRepo),
		frequencyHttp.WithRepository(frequencyRepo),
		frequencyHttp.WithEmployeeFrequencyRepo(employeeFrequencyRepo),
	)

	server.PrintRouter()

	if err := server.Start("0.0.0.0:8081"); err != nil {
		logger.Of(ctx).Fatalf("[main] server.NewServer() returned error: %+v\n", err)
	}
}

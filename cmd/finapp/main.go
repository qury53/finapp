package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/qury53/finapp/internal/core/logger"
	core_repository_pgx "github.com/qury53/finapp/internal/core/repository/postgres"
	core_repository_redis "github.com/qury53/finapp/internal/core/repository/redis"
	transport_http_middleware "github.com/qury53/finapp/internal/core/transport/http/middleware"
	transport_http_server "github.com/qury53/finapp/internal/core/transport/http/server"
	auth_service "github.com/qury53/finapp/internal/features/auth/service"
	auth_transport_http "github.com/qury53/finapp/internal/features/auth/transport/http"
	frontend_repository_filesys "github.com/qury53/finapp/internal/features/frontend/repository/fileSys"
	frontend_service "github.com/qury53/finapp/internal/features/frontend/service"
	frontend_transport_http "github.com/qury53/finapp/internal/features/frontend/transport/http"
	income_repository_postgres "github.com/qury53/finapp/internal/features/income/repository/postgres"
	income_service "github.com/qury53/finapp/internal/features/income/service"
	income_transport_http "github.com/qury53/finapp/internal/features/income/transport/http"
	spending_repository_postgres "github.com/qury53/finapp/internal/features/spendings/repository/postgres"
	spending_service "github.com/qury53/finapp/internal/features/spendings/service"
	spending_transport_http "github.com/qury53/finapp/internal/features/spendings/transport/http"
	statistics_service "github.com/qury53/finapp/internal/features/statistics/service"
	statistics_transport_http "github.com/qury53/finapp/internal/features/statistics/transport/http"
	user_repository_postgres "github.com/qury53/finapp/internal/features/users/repository/postgres"
	users_repository_redis "github.com/qury53/finapp/internal/features/users/repository/redis"
	users_service "github.com/qury53/finapp/internal/features/users/service"
	user_transport_http "github.com/qury53/finapp/internal/features/users/transport/http"
)

// @title Golang FINAPP API
// @version 1.0
// @description FINANCIAL APPLICATION FOR TRACKING YOUR FINANCE
// @contact.name qury53
// @contact.email qury53@yandex.ru
// @BasePath /api/v1
func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	defer cancel()

	//creating logger
	loggerConfig := core_logger.NewConfigMust()
	logger, err := core_logger.NewLogger(loggerConfig)
	if err != nil {
		slog.Error("failed to create logger", "error", err)
		os.Exit(0)
	}

	//creating conn pool
	slog.Debug("creating pool")
	poolConfig := core_repository_pgx.NewConfigMust()
	redisConfig := core_repository_redis.NewConfigMust()
	redisClient, err := core_repository_redis.NewClient(ctx, redisConfig)
	if err != nil {
		logger.Error(
			"failed to craete redisClient",
			slog.String("error", err.Error()),
		)

		os.Exit(0)
	}

	pool, err := core_repository_pgx.NewConnectionPool(ctx, poolConfig)
	if err != nil {
		logger.Error("pool error", "error", err)
	}

	//creating feature users
	logger.Debug("creating users feature")
	usersPostgresRepository := user_repository_postgres.NewUsersRepository(pool)
	usersRedisRepository := users_repository_redis.NewUsersRedisRepository(
		redisClient,
		usersPostgresRepository,
	)

	usersService := users_service.NewUsersService(usersRedisRepository)
	usersHandler := user_transport_http.NewUsersHTTPHandler(usersService)

	//creating feature auth
	logger.Debug("creating auth feature")
	authConfig := auth_service.NewConfigMust()
	authService := auth_service.NewAuthService(authConfig, usersRedisRepository, usersRedisRepository)
	authHandler := auth_transport_http.NewAuthHTTPHandler(authService)

	//creating server
	serverConfig := transport_http_server.NewConfigMust()
	server := transport_http_server.NewHTTPServer(
		serverConfig,
		logger,
		transport_http_middleware.AuthMiddleware(authService),
		transport_http_middleware.CORS(),
		transport_http_middleware.RequestID(),
		transport_http_middleware.Logger(logger),
		transport_http_middleware.Trace(),
		transport_http_middleware.Panic(),
	)

	apiRouterV1 := transport_http_server.NewAPIRouter("v1")

	//creating feature income
	logger.Debug("creating income feature")
	incomeRepository := income_repository_postgres.NewIncomeRepository(pool)
	incomeService := income_service.NewIncomeService(incomeRepository)
	incomeHandler := income_transport_http.NewIncomeHTTPHandler(incomeService)

	//creating feature spending
	logger.Debug("creating spending feature")
	spendingRepository := spending_repository_postgres.NewSpendingRepository(pool)
	spendingService := spending_service.NewSpendingService(spendingRepository)
	spendingHandler := spending_transport_http.NewSpendingHTTPHandler(spendingService)

	//creating feature statistics
	logger.Debug("creating statistics feature")
	statisticsService := statistics_service.NewStatisticsService(
		redisClient,
		incomeRepository,
		spendingRepository,
		usersRedisRepository,
	)

	statisticsTransport := statistics_transport_http.NewStatisticsHTTPHandler(statisticsService)

	//creating feature frontend
	logger.Debug("creating frontend feature")
	frontendRepository := frontend_repository_filesys.NewFrontendRepository()
	frontendService := frontend_service.NewFrontendService(frontendRepository)
	frontendHTTPHandler := frontend_transport_http.NewFrontendHTTPHandler(frontendService)

	apiRouterV1.RegisterRoutes(usersHandler.Routes()...)
	apiRouterV1.RegisterRoutes(incomeHandler.Routes()...)
	apiRouterV1.RegisterRoutes(spendingHandler.Routes()...)
	apiRouterV1.RegisterRoutes(statisticsTransport.Routes()...)
	apiRouterV1.RegisterRoutes(authHandler.Routes()...)
	server.RegisterAPIRouters(*apiRouterV1)
	server.RegisterRoutes(frontendHTTPHandler.Routes()...)
	server.RegisterSwagger()
	if err := server.Run(ctx); err != nil {
		logger.Error("server error", "error", err)
	}
}

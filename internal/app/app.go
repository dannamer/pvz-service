package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dannamer/pvz-service/internal/config"
	"github.com/dannamer/pvz-service/internal/controller/rest"
	"github.com/dannamer/pvz-service/internal/generated/api"
	"github.com/dannamer/pvz-service/internal/infrastructure/jwt"
	"github.com/dannamer/pvz-service/internal/infrastructure/logger"
	"github.com/dannamer/pvz-service/internal/infrastructure/middleware"
	"github.com/dannamer/pvz-service/internal/infrastructure/postgres"
	"github.com/dannamer/pvz-service/internal/repository"
	"github.com/dannamer/pvz-service/internal/usecase"
	"github.com/gin-gonic/gin"
)

func App() {
	zlog := logger.New(true)
	config := config.New()
	ctx := context.Background()
	jwt := jwt.New(config.JwtKey())
	postgres, err := postgres.New(ctx, config.PgUrl())
	if err != nil {
		zlog.With(ctx).Error(ctx, "", err)
	}
	defer postgres.Close()

	repo := repository.New(postgres.Pool)
	uc := usecase.New(jwt, repo)
	handler := rest.NewHandlers(uc, zlog)

	server, err := api.NewServer(handler, middleware.New(jwt))
	if err != nil {
		log.Fatalf("failed to create ogen server: %v", err)
	}

	router := gin.Default()
	router.NoRoute(gin.WrapH(server))
	zlog.Info(ctx, "server successfully started on port 8080")
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.HttpPort()), router); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

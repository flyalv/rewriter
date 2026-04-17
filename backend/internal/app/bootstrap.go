package app

import (
	"backend/internal/config"
	"backend/internal/health"
	"backend/internal/rewrite"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	Router *gin.Engine
	config *config.Config
}

func New() *App {
	cfg := config.Load()

	grpcClient := initGRPC(cfg.GRPCAddress())
	rewriteService := rewrite.NewService(grpcClient)
	rewriteHandler := rewrite.NewHandler(rewriteService)

	router := gin.Default()
	router.GET("/health", health.Check)
	router.POST("/api/v1/rewrite", rewriteHandler.Handle)

	return &App{Router: router, config: cfg}
}

func initGRPC(grpcAddr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to Python gRPC server: %v", err)
	}
	return conn
}

func (a *App) Run() error {
	return a.Router.Run(fmt.Sprintf(":%s", a.config.HTTPPort))
}

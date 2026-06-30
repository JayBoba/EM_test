package main

import (
	"fmt"
	"net/http"

	_ "github.com/JayBoba/EM_test/docs"
	"github.com/JayBoba/EM_test/internal/config"
	"github.com/JayBoba/EM_test/internal/db"
	"github.com/JayBoba/EM_test/internal/handlers"
	"github.com/JayBoba/EM_test/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title Subscription Aggregation API
// @version 1.0
// @description REST сервис для управления подписками и агрегации
// @host localhost:8080
// @BasePath /api/v1
func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("config error: %v", err))
	}
	var logger *zap.Logger
	if cfg.LogLevel == "debug" {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}
	defer logger.Sync()

	dbConn, err := db.NewPostgresDB(cfg.DSN())
	if err != nil {
		logger.Fatal("db connection failed", zap.Error(err))
	}

	r := gin.New()
	r.Use(middleware.GinZap(logger))
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	api.Use(func(c *gin.Context) {
		c.Set("db", dbConn)
		c.Next()
	})

	{
		api.POST("/subscriptions", handlers.CreateSubscription)
		api.GET("/subscriptions/:id", handlers.GetSubscription)
		api.PUT("/subscriptions/:id", handlers.UpdateSubscription)
		api.DELETE("/subscriptions/:id", handlers.DeleteSubscription)
		api.GET("/subscriptions", handlers.ListSubscriptions)
		api.GET("/subscriptions/aggregate", handlers.AggregateCost)
	}
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	logger.Info("starting server", zap.String("addr", addr))
	if err := http.ListenAndServe(addr, r); err != nil {
		logger.Fatal("server failed", zap.Error(err))
	}
}

package middleware

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinZap(logger *zap.Logger) gin.HandlerFunc {
	return ginzap.Ginzap(logger, time.RFC3339, true)
}

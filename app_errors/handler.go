package app_errors

import (
	"github.com/emekarr/coding-test-busha/logger"
	"github.com/emekarr/coding-test-busha/server_response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func ErrorHandler(ctx *gin.Context, err RequestError, fields ...zapcore.Field) {
	logger.Error(err, fields...)
	ctx.Abort()
	server_response.Respond(ctx, err.StatusCode, err.Error(), false, nil)
}

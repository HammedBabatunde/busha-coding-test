package server_response

import (
	"github.com/gin-gonic/gin"
)

func Respond(ctx *gin.Context, code int, message string, success bool, payload interface{}) {
	ctx.JSON(code, gin.H{
		"success": success,
		"message": message,
		"body":    payload,
	})
}

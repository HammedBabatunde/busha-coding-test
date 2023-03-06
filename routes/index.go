package routes

import (
	v1 "github.com/emekarr/coding-test-busha/routes/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter(rg *gin.RouterGroup) {
	v1.InitRoutesV1(rg)
}

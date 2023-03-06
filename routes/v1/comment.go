package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/emekarr/coding-test-busha/app/comment/controller"
)

func InitRoutesV1Comment(rg *gin.RouterGroup) {
	rgComment := rg.Group("/comment")
	{
		rgComment.POST("/create", controller.CreateComment)
	}
}

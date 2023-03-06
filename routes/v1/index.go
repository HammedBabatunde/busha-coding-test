package v1

import "github.com/gin-gonic/gin"

func InitRoutesV1(rg *gin.RouterGroup) {
	rgV1 := rg.Group("/v1")
	{
		InitRoutesV1Comment(rgV1)
	}
}

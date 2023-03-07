package v1

import (
	"github.com/emekarr/coding-test-busha/app/character"
	"github.com/gin-gonic/gin"
)

func InitRoutesV1Characters(rg *gin.RouterGroup) {
	rgComment := rg.Group("/characters")
	{
		rgComment.GET("/fetch", character.FetchAllCharacters)
	}
}

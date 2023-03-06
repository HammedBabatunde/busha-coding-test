package v1

import (
	"github.com/emekarr/coding-test-busha/app/movie"
	"github.com/gin-gonic/gin"
)

func InitRoutesV1Movies(rg *gin.RouterGroup) {
	rgComment := rg.Group("/movies")
	{
		rgComment.GET("/fetch", movie.GetMovies)
	}
}

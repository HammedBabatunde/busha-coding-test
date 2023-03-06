package movie

import (
	"errors"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/emekarr/coding-test-busha/app/comment/models"
	commentRepo "github.com/emekarr/coding-test-busha/app/comment/repository"
	"github.com/emekarr/coding-test-busha/logger"
	"github.com/emekarr/coding-test-busha/server_response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetMovies(ctx *gin.Context) {
	movies, err := MovieService.FetchMovies()
	if err != nil {
		server_response.Respond(ctx, http.StatusInternalServerError, "could not fetch movies", false, nil, nil)
		return
	}
	var moviePayload []ServerResponseMovieType
	var wg sync.WaitGroup
	commentCount := make(chan int64, len(*movies))
	for i, m := range *movies {
		m.FormatCrawl()
		(*movies)[i] = m
		wg.Add(1)
		go func(commentCountChan chan int64, movieName string) {
			commentRepository := commentRepo.GetCommentRepository()
			count, err := commentRepository.CountDocs(models.Comment{Name: movieName})
			if err != nil {
				logger.Error(errors.New("could not count movie docs"), zap.Error(err))
				commentCountChan <- 0
				return
			}
			commentCountChan <- *count
		}(commentCount, m.Title)
		moviePayload = append(moviePayload, ServerResponseMovieType{
			CommentCount: <-commentCount,
			Movie:        m,
		})
	}
	sort.Slice(moviePayload, func(i, j int) bool {
		k, err := time.Parse("2006-01-02", moviePayload[i].RelaseDate)
		if err != nil {
			logger.Error(errors.New("failed to convert release date to time.time for fetch movies"), zap.Error(err), zap.String("time", moviePayload[i].RelaseDate))
		}
		l, err := time.Parse("2006-01-02", moviePayload[j].RelaseDate)
		if err != nil {
			logger.Error(errors.New("failed to convert release date to time.time for fetch movies"), zap.Error(err), zap.String("time", moviePayload[j].RelaseDate))
		}
		return k.Before(l)
	})
	server_response.Respond(ctx, http.StatusOK, "movies fetched", true, moviePayload, nil)
}

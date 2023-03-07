package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/emekarr/coding-test-busha/app/comment"
	"github.com/emekarr/coding-test-busha/app/comment/models"
	comentRepo "github.com/emekarr/coding-test-busha/app/comment/repository"
	commentUseCases "github.com/emekarr/coding-test-busha/app/comment/usecases/comment"
	"github.com/emekarr/coding-test-busha/app/movie"
	"github.com/emekarr/coding-test-busha/app_errors"
	"github.com/emekarr/coding-test-busha/server_response"
	"github.com/gin-gonic/gin"
)

func CreateComment(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		server_response.Respond(ctx, http.StatusBadRequest, "pass in a search term", false, nil, nil)
		return
	}
	var commentPayload comment.CommentPayload
	if err := ctx.ShouldBindJSON(&commentPayload); err != nil {
		server_response.Respond(ctx, http.StatusBadRequest, "pass in a valid json object that matches the payload", false, nil, nil)
		return
	}
	if errs := commentPayload.Validate(); errs != nil {
		server_response.Respond(ctx, http.StatusBadRequest, "passing in the correct comment payload", false, nil, errs)
		return
	}
	moviePayload, err := movie.MovieService.SearchMovies(id)
	if err != nil {
		app_errors.ErrorHandler(ctx, app_errors.RequestError{Err: errors.New("an error occured while fetching movies"), StatusCode: http.StatusInternalServerError})
		return
	}
	if (moviePayload == nil || *moviePayload == movie.Movie{}) {
		app_errors.ErrorHandler(ctx, app_errors.RequestError{Err: fmt.Errorf("movie with id=%s does not exist", id), StatusCode: http.StatusInternalServerError})
		return
	}
	result, errs := commentUseCases.CreateCommentUseCase(ctx, models.Comment{
		Name:    moviePayload.Title,
		Comment: commentPayload.Comment,
	})
	if errs != nil {
		server_response.Respond(ctx, http.StatusBadRequest, "could not save comment", false, result, errs)
		return
	}
	server_response.Respond(ctx, http.StatusCreated, "comment saved", true, result, nil)
}

func FetchComments(ctx *gin.Context) {
	var movieName struct {
		Name string `json:"name"`
	}
	if err := ctx.ShouldBindJSON(&movieName); err != nil {
		server_response.Respond(ctx, http.StatusBadRequest, "pass in a valid json object that matches the payload", false, nil, nil)
		return
	}
	commentRepository := comentRepo.GetCommentRepository()
	comments, err := commentRepository.RunRawSQLFind("SELECT * FROM comments WHERE comments.name LIKE ? ORDER BY comments.created_at DESC;", movieName.Name+"%")
	if err != nil {
		app_errors.ErrorHandler(ctx, app_errors.RequestError{Err: err, StatusCode: http.StatusInternalServerError})
		return
	}
	server_response.Respond(ctx, http.StatusOK, "comments fetched", true, comments, nil)
}

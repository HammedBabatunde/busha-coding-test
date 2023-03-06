package controller

import (
	"errors"
	"net/http"

	"github.com/emekarr/coding-test-busha/app/comment"
	"github.com/emekarr/coding-test-busha/app/comment/models"
	commentUseCases "github.com/emekarr/coding-test-busha/app/comment/usecases/comment"
	"github.com/emekarr/coding-test-busha/app/movie"
	"github.com/emekarr/coding-test-busha/app_errors"
	"github.com/emekarr/coding-test-busha/server_response"
	"github.com/gin-gonic/gin"
)

func CreateComment(ctx *gin.Context) {
	term := ctx.Query("term")
	if term == "" {
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
	movie, err := movie.MovieService.SearchMovies(term)
	if err != nil {
		app_errors.ErrorHandler(ctx, app_errors.RequestError{Err: errors.New("pass in a json object"), StatusCode: http.StatusInternalServerError})
		return
	}
	result, errs := commentUseCases.CreateCommentUseCase(ctx, models.Comment{
		Name:    movie.Title,
		Comment: commentPayload.Comment,
	})
	if errs != nil {
		server_response.Respond(ctx, http.StatusBadRequest, "could not save comment", false, result, errs)
		return
	}
	server_response.Respond(ctx, http.StatusCreated, "comment saved", true, result, nil)
}

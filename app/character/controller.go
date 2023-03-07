package character

import (
	"net/http"

	"github.com/emekarr/coding-test-busha/app_errors"
	"github.com/emekarr/coding-test-busha/server_response"
	"github.com/gin-gonic/gin"
)

func FetchAllCharacters(ctx *gin.Context) {
	name := ctx.Query("name")
	gender := ctx.Query("gender")
	height := ctx.Query("height")
	order := ctx.Query("asc")
	filter := ctx.Query("filter")
	page := ctx.Query("page")

	characters, err := CharacterService.FetchAllMovieCharacters(page)
	if err != nil {
		app_errors.ErrorHandler(ctx, app_errors.RequestError{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		})
		return
	}
	characters = CharacterService.FilterMovieCharacters(characters, filter)
	if name == "true" {
		CharacterService.SortMovieCharacters(characters, "name", order == "true")
	} else if gender == "true" {
		CharacterService.SortMovieCharacters(characters, "gender", order == "true")
	} else if height == "true" {
		CharacterService.SortMovieCharacters(characters, "height", order == "true")
	} else {
		server_response.Respond(ctx, http.StatusBadRequest, "no sort paramter added", false, nil, &[]string{"pass in a sort parameter"})
		return
	}
	parsedCharacters := CharacterService.SetCharacterHeightValueFeet(characters)
	totalHeight := CharacterService.ClculateTotalCharacterHeightCM(parsedCharacters)
	server_response.Respond(ctx, http.StatusOK, "characters fetched", true, map[string]interface{}{
		"characters":    parsedCharacters,
		"count":         len(*parsedCharacters),
		"totalHeigthCM": totalHeight,
	}, nil)
}

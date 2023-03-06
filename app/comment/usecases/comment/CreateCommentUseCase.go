package comment

import (
	"github.com/emekarr/coding-test-busha/app/comment/models"
	commentRepo "github.com/emekarr/coding-test-busha/app/comment/repository"
	"github.com/gin-gonic/gin"
)

func CreateCommentUseCase(ctx *gin.Context, payload models.Comment) (*models.Comment, *[]string) {
	payload.CommenterIP = ctx.ClientIP()
	payload.InitFields()
	if errs := payload.Validate(); errs != nil {
		return nil, errs
	}
	commentRepository := commentRepo.GetCommentRepository()
	result, err := commentRepository.CreateOne(&payload)
	if err != nil {
		return nil, &[]string{err.Error()}
	}
	return result, nil
}

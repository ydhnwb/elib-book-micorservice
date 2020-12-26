package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/elib-book-microservice/application/repository"
	"github.com/ydhnwb/elib-book-microservice/domain/entity"
	"github.com/ydhnwb/elib-book-microservice/infrastructure/helper"
)

//BookDetailUseCase is a contract
type BookDetailUseCase interface {
	DetailBook(ctx *gin.Context)
}

type bookDetailUseCase struct {
	bookRepository repository.BookRepository
}

//NewBookDetailUseCase creates a new instance of BookDetailUseCase
func NewBookDetailUseCase(repo repository.BookRepository) BookDetailUseCase {
	return &bookDetailUseCase{
		bookRepository: repo,
	}
}

func (ctl *bookDetailUseCase) DetailBook(ctx *gin.Context) {
	id := ctx.Param("id")
	book := entity.Book{}
	book = ctl.bookRepository.FindBookByID(id)
	helper.BuildResponse(http.StatusOK, book, ctx)
}

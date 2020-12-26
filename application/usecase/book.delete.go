package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/elib-book-microservice/application/repository"
	"github.com/ydhnwb/elib-book-microservice/infrastructure/helper"
)

//BookDeleteUseCase is a contract
type BookDeleteUseCase interface {
	DeleteBook(ctx *gin.Context)
}

type bookDeleteUseCase struct {
	bookRepository repository.BookRepository
}

//NewBookDeleteUseCase creates a new instance BookDeleteUseCase
func NewBookDeleteUseCase(repo repository.BookRepository) BookDeleteUseCase {
	return &bookDeleteUseCase{
		bookRepository: repo,
	}
}

func (ctl *bookDeleteUseCase) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	book := ctl.bookRepository.FindBookByID(id)
	ctl.bookRepository.DeleteBook(book)
	helper.BuildResponse(http.StatusOK, helper.EmptyObj{}, ctx)
}

package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/elib-book-microservice/application/repository"
	"github.com/ydhnwb/elib-book-microservice/infrastructure/helper"
)

//BookListUseCase is a contract
type BookListUseCase interface {
	BookList(ctx *gin.Context)
}

type bookListUseCase struct {
	bookRepository repository.BookRepository
}

//NewBookListUseCase creates a new instance of BookListUseCase
func NewBookListUseCase(repo repository.BookRepository) BookListUseCase {
	return &bookListUseCase{
		bookRepository: repo,
	}
}

func (ctl *bookListUseCase) BookList(ctx *gin.Context) {
	books := ctl.bookRepository.AllBook()
	helper.BuildResponse(http.StatusOK, books, ctx)
}

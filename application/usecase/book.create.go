package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/elib-book-microservice/application/repository"
)

//BookCreateUseCase is a contract
type BookCreateUseCase interface {
	CreateBook(ctx *gin.Context)
}

type bookCreateUseCase struct {
	bookRepository repository.BookRepository
}

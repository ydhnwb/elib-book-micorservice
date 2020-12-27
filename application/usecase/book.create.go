package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/ydhnwb/elib-book-microservice/application/repository"
	"github.com/ydhnwb/elib-book-microservice/domain/entity"
	"github.com/ydhnwb/elib-book-microservice/infrastructure/dto"
	"github.com/ydhnwb/elib-book-microservice/infrastructure/helper"
)

//BookCreateUseCase is a contract
type BookCreateUseCase interface {
	CreateBook(ctx *gin.Context)
}

type bookCreateUseCase struct {
	bookRepository repository.BookRepository
}

//NewBookCreateUseCase creates a new instance of BookCreateUseCase
func NewBookCreateUseCase(repo repository.BookRepository) BookCreateUseCase {
	return &bookCreateUseCase{
		bookRepository: repo,
	}
}

func (ctl *bookCreateUseCase) CreateBook(ctx *gin.Context) {
	bookCreateDTO := dto.BookCreateDTO{}
	e := ctx.ShouldBind(&bookCreateDTO)
	if e != nil {
		helper.BuildErrorResponse(http.StatusBadRequest, e.Error(), helper.EmptyObj{}, ctx)
		return
	}
	book := entity.Book{}
	e = smapping.FillStruct(&book, smapping.MapFields(&bookCreateDTO))
	if e != nil {
		helper.BuildErrorResponse(http.StatusBadRequest, e.Error(), helper.EmptyObj{}, ctx)
		return
	}
	result := ctl.bookRepository.InsertBook(book)
	helper.BuildResponse(http.StatusCreated, result, ctx)

}

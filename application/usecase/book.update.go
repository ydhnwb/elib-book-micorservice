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

//BookUpdateUseCase is a contract
type BookUpdateUseCase interface {
	BookUpdate(ctx *gin.Context)
}

type bookUpdateUseCase struct {
	bookRepository repository.BookRepository
}

//NewBookUpdateUseCase creates a new instance of BookUpdateUseCase
func NewBookUpdateUseCase(repo repository.BookRepository) BookUpdateUseCase {
	return &bookUpdateUseCase{
		bookRepository: repo,
	}
}

func (ctl *bookUpdateUseCase) BookUpdate(ctx *gin.Context) {
	bookUpdateDTO := dto.BookUpdateDTO{}
	e := ctx.ShouldBind(&bookUpdateDTO)
	if e != nil {
		helper.BuildErrorResponse(http.StatusBadRequest, e.Error(), helper.EmptyObj{}, ctx)
		return
	}

	book := entity.Book{}
	e = smapping.FillStruct(&book, smapping.MapFields(&bookUpdateDTO))
	if e != nil {
		helper.BuildErrorResponse(http.StatusBadRequest, e.Error(), helper.EmptyObj{}, ctx)
		return
	}

	ctl.bookRepository.UpdateBook(book)
	helper.BuildResponse(http.StatusOK, book, ctx)

}

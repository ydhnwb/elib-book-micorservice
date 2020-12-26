package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/elib-book-microservice/application/repository"
	"github.com/ydhnwb/elib-book-microservice/application/usecase"
	"github.com/ydhnwb/elib-book-microservice/infrastructure/persistence"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                  = persistence.SetupDatabaseConnection()
	bookRepository    repository.BookRepository = repository.NewBookRepository(db)
	bookListUseCase   usecase.BookListUseCase   = usecase.NewBookListUseCase(bookRepository)
	bookCreateUseCase usecase.BookCreateUseCase = usecase.NewBookCreateUseCase(bookRepository)
	bookUpdateUseCase usecase.BookUpdateUseCase = usecase.NewBookUpdateUseCase(bookRepository)
	bookDeleteUseCase usecase.BookDeleteUseCase = usecase.NewBookDeleteUseCase(bookRepository)
	bookDetailUseCase usecase.BookDetailUseCase = usecase.NewBookDetailUseCase(bookRepository)
)

func main() {
	defer persistence.CloseDatabaseConnection(db)
	r := gin.Default()

	bookRoutes := r.Group("api/books")
	{
		bookRoutes.GET("/", bookListUseCase.BookList)
		bookRoutes.POST("/", bookCreateUseCase.CreateBook)
		bookRoutes.GET("/:id", bookDetailUseCase.DetailBook)
		bookRoutes.PUT("/:id", bookUpdateUseCase.BookUpdate)
		bookRoutes.DELETE("/:id", bookDeleteUseCase.DeleteBook)
	}

	r.Run(":8082")

}

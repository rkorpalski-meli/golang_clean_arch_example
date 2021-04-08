package registry

import (
	"github.com/rkorpalski-meli/golang_clean_arch_example/src/api/application/handlers"
	"github.com/rkorpalski-meli/golang_clean_arch_example/src/api/entity"
	"github.com/rkorpalski-meli/golang_clean_arch_example/src/api/infrastructre/repository"
	"github.com/rkorpalski-meli/golang_clean_arch_example/src/api/usecase"
)

func CreateBookHandler() handlers.BookHandler  {
	return handlers.NewBookHandler(createBookUsecase())
}

func createBookUsecase() entity.BookUsecase {
	return usecase.NewBookUsecase(createBookRepository())
}

func createBookRepository() entity.BookRepository {
	return repository.NewBookCacheRepository()
}

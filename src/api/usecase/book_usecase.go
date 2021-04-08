package usecase

import (
	"errors"
	"github.com/rkorpalski-meli/golang_clean_arch_example/src/api/entity"
)

type bookUsecase struct {
	bookRepository entity.BookRepository
}

func NewBookUsecase(bookRepository entity.BookRepository) entity.BookUsecase {
	return bookUsecase{bookRepository: bookRepository}
}

func(u bookUsecase) FindBook(title string) (entity.Book, error) {
	book, err := u.bookRepository.Get(title)
	if err != nil &&  errors.Is(err, entity.ErrBookNotFound) {
		return book, nil
	}

	return book, err
}

func(u bookUsecase) CreateBook(title string, author string) error {
	newBook := entity.Book{
		Title:  title,
		Author: author,
	}

	if  isValid := newBook.Validate(); !isValid {
		return entity.ErrBookNotValid
	}

	return u.bookRepository.Insert(entity.Book{
		Title:  title,
		Author: author,
	})
}
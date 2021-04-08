package repository

import "github.com/rkorpalski-meli/golang_clean_arch_example/src/api/entity"

type BookCacheRepository struct {
	cache 	[]entity.Book
}

func NewBookCacheRepository() entity.BookRepository {
	return BookCacheRepository{cache: make([]entity.Book, 0)}
}

func (r BookCacheRepository) Get(title string) (entity.Book, error) {
	for _, book := range r.cache {
		if book.Title == title {
			return book, nil
		}
	}
	return entity.Book{}, entity.ErrBookNotFound
}

func(r BookCacheRepository) Insert(book entity.Book) error {
	r.cache = append(r.cache, book)
	return nil
}
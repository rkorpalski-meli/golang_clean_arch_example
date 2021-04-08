package usecase

import (
	"github.com/rkorpalski-meli/golang_clean_arch_example/src/api/entity"
	"github.com/rkorpalski-meli/golang_clean_arch_example/src/api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestBookUsecase_FindBook(t *testing.T) {
	testCases := []struct{
		testName	string
		bookTitle	string
		bookRepository	func(mock *mocks.BookRepository)
		expectedResult	entity.Book
		expectedError	error
	}{
		{
			testName:  "Find a valid book",
			bookTitle: "The Lord of the Rings",
			bookRepository: func(m *mocks.BookRepository) {
				m.On("Get", mock.Anything).Return(entity.Book{
					Title:  "The Lord of the Rings",
					Author: "J.R.R Tolkien",
				}, nil)
			},
			expectedResult: entity.Book{
				Title:  "The Lord of the Rings",
				Author: "J.R.R Tolkien",
			},
			expectedError:  nil,
		},
		{
			testName:  "Book not found",
			bookTitle: "The Divine Comedy",
			bookRepository: func(m *mocks.BookRepository) {
				m.On("Get", mock.Anything).Return(entity.Book{}, entity.ErrBookNotFound)
			},
			expectedResult: entity.Book{},
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			bookRepository := new(mocks.BookRepository)
			tc.bookRepository(bookRepository)

			bookUsecase := NewBookUsecase(bookRepository)
			result, err := bookUsecase.FindBook(tc.bookTitle)

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
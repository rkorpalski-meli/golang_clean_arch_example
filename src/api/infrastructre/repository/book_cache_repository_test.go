package repository

import (
	"github.com/rkorpalski-meli/golang_clean_arch_example/src/api/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookCacheRepository_Get(t *testing.T) {
	bookCacheRepository := NewBookCacheRepository()
	bookCacheRepository.Insert(entity.Book{
		Title:  "The Lord of the Rings",
		Author: "J.R.R Tolkien",
	})
	testCases := []struct{
		testName		string
		bookTitle		string
		expectedResult	entity.Book
		expectedError	error
	}{
		{
			testName: "Get a valid Book",
			bookTitle: "The Lord of the Rings",
			expectedResult: entity.Book{
				Title:  "The Lord of the Rings",
				Author: "J.R.R Tolkien",
			},
			expectedError: nil,

		},
		{
			testName: "Book not found",
			bookTitle: "The Divine Comedy",
			expectedResult: entity.Book{},
			expectedError: entity.ErrBookNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result, err := bookCacheRepository.Get(tc.bookTitle)
			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
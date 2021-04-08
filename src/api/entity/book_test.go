package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Book_Validate(t *testing.T) {
	testCases := []struct{
		testName		string
		bookTitle 		string
		bookAuthor 		string
		expectedResult	bool
	}{
		{
			testName: "Valid Book",
			bookTitle: "The Lord of the Rings",
			bookAuthor: "J.R.R Tolkien",
			expectedResult: true,
		},
		{
			testName: "Invalid Book - Title is empty",
			bookTitle: "",
			bookAuthor: "J.R.R Tolkien",
			expectedResult: false,
		},
		{
			testName: "Invalid Book - Author is empty",
			bookTitle: "The Lord of the Rings",
			bookAuthor: "",
			expectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			book := Book{
				Title:  tc.bookTitle,
				Author: tc.bookAuthor,
			}
			result := book.Validate()
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

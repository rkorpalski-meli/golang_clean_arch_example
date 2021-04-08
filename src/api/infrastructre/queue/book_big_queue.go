package queue

import (
	"fmt"
	"github.com/rkorpalski-meli/golang_clean_arch_example/src/api/entity"
)

type BookBigQueue struct{}

func NewBookBigQueue() entity.BookQueue {
	return BookBigQueue{}
}

func(q BookBigQueue) SendToTopic(book entity.Book) error {
	fmt.Println("Sending message to BigQueue Topic")
	return nil
}

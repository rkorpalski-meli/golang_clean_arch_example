package entity

type Book struct {
	Title 	string
	Author	string
}

func(b Book) Validate() bool {
	if b.Title == "" || b.Author == "" {
		return false
	}
	return true
}

type BookRepository interface {
	Get(title string) (Book, error)
	Insert(book Book) error
}

type BookUsecase interface {
	FindBook(title string) (Book, error)
	CreateBook(title string, author string) error
}

type BookQueue interface {
	SendToTopic(book Book) error
}
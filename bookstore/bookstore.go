package bookstore

type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(book Book) Book {
	book.Copies--
	return book
}

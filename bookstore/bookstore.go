package bookstore

import "errors"

type Book struct {
	ID     int
	Title  string
	Author string
	Copies int
}

func Buy(book Book) (Book, error) {
	if book.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	book.Copies--
	return book, nil
}

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

func GetBook(catalog []Book, ID int) Book {
	for _, b := range catalog {
		if b.ID == ID {
			return b
		}
	}
	return Book{}
}

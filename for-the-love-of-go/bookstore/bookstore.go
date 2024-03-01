package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	ID              int
	Title           string
	Author          string
	Copies          int
	PriceCents      int
	DiscountPercent int
	category        Category
}
type Category int
type Catalog map[int]Book

const (
	CategoryAutobiography Category = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

var validCategory = map[Category]bool{
	CategoryAutobiography:     true,
	CategoryLargePrintRomance: true, CategoryParticlePhysics: true,
}

func Buy(book Book) (Book, error) {
	if book.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	book.Copies--
	return book, nil
}

func (catalog Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, b := range catalog {
		result = append(result, b)
	}
	return result
}

func (catalog Catalog) GetBook(ID int) (Book, error) {
	book, ok := catalog[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return book, nil
}

func (b Book) NetPriceCents() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("negtaive price %d", price)
	}
	b.PriceCents = price
	return nil
}

func (b Book) Category() Category {
	return b.category
}

func (b *Book) SetCategory(category Category) error {
	if !validCategory[category] {
		return fmt.Errorf("unknown category %v", category)
	}
	b.category = category
	return nil
}

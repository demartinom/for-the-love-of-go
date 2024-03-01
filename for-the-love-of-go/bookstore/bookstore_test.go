package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{ID: 1, Title: "Spark Joy", Author: "Marie Kondo", Copies: 2}
	want := 1
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestBuyErrors(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{ID: 1, Title: "Spark Joy", Author: "Marie Kondo", Copies: 0}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}
	got := catalog.GetAllBooks()
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"}}
	want := bookstore.Book{
		ID: 2, Title: "The Power of Go: Tools"}
	got, err := catalog.GetBook(2)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadIDReturnsError(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{}
	_, err := catalog.GetBook(999)
	if err == nil {
		t.Fatal("want error for non-existend ID, got nil")
	}
}
func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title: "Title", PriceCents: 2000, DiscountPercent: 50}
	want := 1000
	got := book.NetPriceCents()
	if want != got {
		t.Errorf("wanted %d, got %d", want, got)
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{PriceCents: 300}
	want := 200
	err := book.SetPriceCents(200)
	if err != nil {
		t.Fatal(err)
	}
	got := book.PriceCents
	if want != got {
		t.Errorf("want updated price %d, got %d", want, got)
	}
}

func TestInvalidPriceCents(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{PriceCents: 300}
	err := book.SetPriceCents(-1)
	if err == nil {
		t.Fatal("want error setting invalid price -1, got nil")
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{Title: "For the Love of Go"}
	testCases := []struct {
		desc string
		want bookstore.Category
	}{
		{
			desc: "want 0",
			want: bookstore.CategoryAutobiography,
		},
		{desc: "want 1", want: bookstore.CategoryLargePrintRomance}, {desc: "want 2", want: bookstore.CategoryParticlePhysics}}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := book.SetCategory(2)
			if err != nil {
				t.Fatal(err)
			}
			got := book.Category()
			if tC.want != got {
				t.Errorf("want category %v, got %v", tC.want, got)
			}
		})
	}
}

func TestInvalidSetCategory(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{}
	err := book.SetCategory(999)
	if err == nil {
		t.Fatal("want error for invalid category, got nil")
	}
}

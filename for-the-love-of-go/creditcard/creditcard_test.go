package creditcard_test

import (
	"creditcard"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()
	want := "1234567890123456"
	cc, err := creditcard.New(want)
	if err != nil {
		t.Fatal(err)
	}
	got := cc.Number()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestInvalidReturnsError(t *testing.T) {
	t.Parallel()
	_, err := creditcard.New("")
	if err == nil {
		t.Error("wanted error for invalid card number, got nil")
	}
}

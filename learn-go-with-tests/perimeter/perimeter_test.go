package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if want != got {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{3.0, 5.0}
		got := rectangle.Area()
		want := 15.0
		if want != got {
			t.Errorf("got %.g, want %.g", want, got)
		}
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	})
}

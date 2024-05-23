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

// func TestArea(t *testing.T) {
// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g, want %g", got, want)
// 		}
// 	}
// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}
// 		got := rectangle.Area()
// 		checkArea(t, rectangle, got)
// 	})
// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		got := circle.Area()
// 		checkArea(t, circle, got)
// 	})
// }

func TestArea(t *testing.T) {
	testCases := []struct {
		desc  string
		shape Shape
		want  float64
	}{
		{
			"rectangles",
			Rectangle{12, 6},
			72.0,
		},
		{
			"circles",
			Circle{10},
			314.1592653589793,
		},
		{
			"triangles",
			Triangle{12, 6},
			36.0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.shape.Area()
			if got != tC.want {
				t.Errorf("got %g, want %g", got, tC.want)
			}
		})
	}
}

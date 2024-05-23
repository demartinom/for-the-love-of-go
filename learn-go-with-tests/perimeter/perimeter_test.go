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
	testCases := []struct {
		desc  string
		shape Shape
		want  float64
	}{
		{
			desc:  "rectangles",
			shape: Rectangle{12, 6},
			want:  72.0,
		},
		{
			desc:  "circles",
			shape: Circle{10},
			want:  314.1592653589793,
		},
		{
			desc:  "triangles",
			shape: Triangle{12, 6},
			want:  36.0,
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

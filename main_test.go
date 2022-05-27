package square

import "testing"

func TestEnd(t *testing.T) {
	square := Square{Point{1, 1}, 5}
	endPointEtalon := Point{6, 6}
	endPoint := square.End()
	if endPoint != endPointEtalon {
		t.Errorf("Test method End() failed. Expected x: 6, y: 6, but was x: %d, y: %d", endPoint.x, endPoint.y)
	}
}

func TestArea(t *testing.T) {
	square := Square{Point{1, 2}, 5}
	area := square.Area()
	if area != 25 {
		t.Errorf("Test method End() failed. Expected 25, but was %d", area)
	}
}

func TestPerimeter(t *testing.T) {
	square := Square{Point{1, 2}, 5}
	perimeter := square.Perimeter()
	if perimeter != 20 {
		t.Errorf("Test method End() failed. Expected 25, but was %d", perimeter)
	}
}

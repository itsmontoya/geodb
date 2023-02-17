package math

import "testing"
import "fmt"

const testInvalidValueFmt = "invalid value, %v was expected and %x was received\n"

func TestSegment(t *testing.T) {
	s := NewSegment(0, 7, 4, 23)
	if val := s.GetX(23); val != 4 {
		t.Fatalf(testInvalidValueFmt, 4, val)
	}

	if val := s.GetY(4); val != 23 {
		t.Fatalf(testInvalidValueFmt, 23, val)
	}

	if val := s.GetX(7); val != 0 {
		t.Fatalf(testInvalidValueFmt, 0, val)
	}

	if val := s.GetY(0); val != 7 {
		t.Fatalf(testInvalidValueFmt, 7, val)
	}

	if !s.Intersects(4, 23) {
		t.Fatal("false negative for intersection check")
	}

	if !s.Intersects(0, 15) {
		t.Fatal("false negative for intersection check")
	}

	if s.Intersects(-3, 6) {
		t.Fatal("false positive for intersection check")
	}

	fmt.Println("")

	s = NewSegment(0, 7, 4, 7)
	fmt.Println(s.GetY(9))
	fmt.Println(s.GetY(36))
	fmt.Println(s.GetX(7))

	fmt.Println("")

	s = NewSegment(3, 0, 3, 12)
	fmt.Println(s.GetX(0))
	fmt.Println(s.GetX(12))
	fmt.Println(s.GetY(3))
}

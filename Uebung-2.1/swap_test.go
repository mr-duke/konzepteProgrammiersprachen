package main
import "testing"

func TestSwap (t *testing.T) {
	a,b := 1,2 
	a, b = swap(a,b)

	if a != 2 || b != 1 {
		t.Error("Numbers weren't swapped")
	}

}
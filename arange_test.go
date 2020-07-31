package arange

import (
	"testing"
)

/*
example:

arange(0., 10., 0.5)
should return
[0 0.5 1 1.5 2 2.5 3 3.5 4 4.5 5 5.5 6 6.5 7 7.5 8 8.5 9 9.5]
*/

var expectedOutput = []float64{0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5, 9, 9.5}

func TestArange(t *testing.T) {

	uc := arange(0., 10., 0.5)
	if !testEq(uc, expectedOutput, t) {
		t.Errorf("output array not equal to the one expected")
	}
}

// compare if two arrays are equal
func testEq(a, b []float64, t *testing.T) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

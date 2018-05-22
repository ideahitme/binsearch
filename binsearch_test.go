package binsearch

import "testing"

var predNotNegative Pred = Pred(func(el int) bool {
	// returns true when element is non-negative
	return el > -1
})

func TestFirstTrue(t *testing.T) {
	t.Run("Test FirstTrue for array of size: 1", testFirstTrueSizeOne)
	t.Run("Test FirstTrue for array of size: 2", testFirstTrueSizeTwo)
	t.Run("Test FirstTrue for array of size more than 1", testFirstTrueLarge)
}

func TestLastFalse(t *testing.T) {
	t.Run("Test LastFalse for array of size: 1", testLastFalseSizeOne)
	t.Run("Test LastFalse for array of size: 2", testLastFalseSizeTwo)
	t.Run("Test LastFalse for array of size more than 1", testLastFalseLarge)
}

func testFirstTrueSizeOne(t *testing.T) {
	x := []int{1}
	index, err := FirstTrue(x, predNotNegative)
	if err != nil {
		t.Fatal("No error is expected")
	}
	if index != 0 {
		t.Error("Should return first element because it is non-negative")
	}

	y := []int{-1}
	index, err = FirstTrue(y, predNotNegative)
	if index != 0 {
		t.Error("No match is expected. Index is set to zero")
	}
	if err != NoMatch {
		t.Error("Expected no match error")
	}
}

func testFirstTrueSizeTwo(t *testing.T) {
	x := []int{-1, 1}
	index, err := FirstTrue(x, predNotNegative)
	if err != nil {
		t.Fatal("No error is expected")
	}
	if index != 1 {
		t.Error("Should return first element because it is non-negative")
	}

	y := []int{-2, -1}
	index, err = FirstTrue(y, predNotNegative)
	if index != 0 {
		t.Error("No match is expected. Index is set to zero")
	}
	if err != NoMatch {
		t.Error("Expected no match error")
	}
}

func testFirstTrueLarge(t *testing.T) {
	x := []int{0, 1, 2, 3, 4}
	index, err := FirstTrue(x, predNotNegative)
	if err != nil {
		t.Fatal("No error is expected")
	}
	if index != 0 {
		t.Error("Should return correct index")
	}

	x = []int{-2, -1, 2, 3, 4}
	index, err = FirstTrue(x, predNotNegative)
	if err != nil {
		t.Fatal("No error is expected")
	}
	if index != 2 {
		t.Error("Should return correct index")
	}

	x = []int{-10, -5, -4, 0}
	index, err = FirstTrue(x, predNotNegative)
	if err != nil {
		t.Fatal("No error is expected")
	}
	if index != 3 {
		t.Error("Should return correct index")
	}

	x = []int{-10, -5, -4, -1}
	index, err = FirstTrue(x, predNotNegative)
	if err != NoMatch {
		t.Fatal("NoMatch error is expected")
	}
	if index != 0 {
		t.Error("Should return correct index")
	}

}
func testLastFalseSizeOne(t *testing.T) {
	x := []int{1}
	index, err := LastFalse(x, predNotNegative)
	if err != NoMatch {
		t.Fatal("NoMatch error is expected")
	}
	if index != 0 {
		t.Error("Should return first element because it is non-negative")
	}

	y := []int{-1}
	index, err = LastFalse(y, predNotNegative)
	if index != 0 {
		t.Error("No match is expected. Index is set to zero")
	}
	if err != nil {
		t.Error("No error is expected")
	}
}

func testLastFalseSizeTwo(t *testing.T) {
	x := []int{-1, 1}
	index, err := LastFalse(x, predNotNegative)
	if err != nil {
		t.Fatal("No error is expected")
	}
	if index != 0 {
		t.Error("Should return first element because it is last negative")
	}

	y := []int{-2, -1}
	index, err = LastFalse(y, predNotNegative)
	if index != 1 {
		t.Error("Should return last negative element")
	}
	if err != nil {
		t.Error("Expected no match error")
	}

	z := []int{1, 2}
	index, err = LastFalse(z, predNotNegative)
	if index != 0 {
		t.Error("Should return 0 when no match is found")
	}
	if err != NoMatch {
		t.Error("NoMatch error is expected")
	}
}

func testLastFalseLarge(t *testing.T) {
	x := []int{0, 1, 2, 3, 4}
	index, err := LastFalse(x, predNotNegative)
	if err != NoMatch {
		t.Fatal("NoMatch error is expected")
	}
	if index != 0 {
		t.Error("Should return correct index")
	}

	x = []int{-2, -1, 2, 3, 4}
	index, err = LastFalse(x, predNotNegative)
	if err != nil {
		t.Fatal("No error is expected")
	}
	if index != 1 {
		t.Error("Should return correct index")
	}

	x = []int{-10, -5, -4, 0}
	index, err = LastFalse(x, predNotNegative)
	if err != nil {
		t.Fatal("No error is expected")
	}
	if index != 2 {
		t.Error("Should return correct index")
	}

	x = []int{-10, -5, -4, -1}
	index, err = LastFalse(x, predNotNegative)
	if err != nil {
		t.Fatal("No error is expected")
	}
	if index != 3 {
		t.Error("Should return correct index")
	}

}

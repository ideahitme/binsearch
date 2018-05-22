### Binary search 

Very simple library for binary search in Go. Only works with arrays of `int` for any custom predicate (type defined)

Usage:

```go

import ( 
	"fmt" 

	"github.com/ideahitme/binsearch"
)

func main() { 
	// pred returns true if element is nonnegative
	pred := binsearch.Pred(func(el int) bool { 
		return el > -1 
	})

	// find index of first non-negative element in the sorted array: 
	x := []int{ -5, -3, 1, 5 }
	index, err := binsearch.FirstTrue(x, pred)
	fmt.Println(index, err) // 2 nil

	// if no match error is returned	
	y := []int{ -5, -3, -1 }
	index, err = binsearch.FirstTrue(y, pred)
	fmt.Println(index, err) // 0 No match is found

	// find index of last element for which predicate is false
	index, err = binsearch.LastFalse(x, pred)
	fmt.Println(index, err) // 1, nil
}

```

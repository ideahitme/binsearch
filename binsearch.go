package binsearch

import "fmt"

var (
	NoMatch error = fmt.Errorf("No match is found")
)

/**
In all provided APIs arrays are assumed to be sorted according to the predicate `p`, i.e.
for any i < j, if p(x[i]) is true then p(x[j]) is true
*/

// Pred is a predicate function, it accepts the element as a parameter and returns true
// if the condition is satisfied
type Pred func(el int) bool

// FirstTrue returns first index such that predicate is true
// if no such element NoMatch error is returned
func FirstTrue(x []int, p Pred) (int, error) {
	l := 0
	r := len(x) - 1

	for l < r {
		m := l + (r-l)/2
		if p(x[m]) {
			r = m
		} else {
			l = m + 1
		}
	}
	if p(x[l]) {
		return l, nil
	}
	return 0, NoMatch
}

// LastFalse returns last index such taht predicate is true
// if no such element NoMatch error is returned
func LastFalse(x []int, p Pred) (int, error) {
	l := 0
	r := len(x) - 1
	for l < r {
		m := l + (r-l+1)/2
		if p(x[m]) {
			r = m - 1
		} else {
			l = m
		}
	}
	if !p(x[l]) {
		return l, nil
	}
	return 0, NoMatch
}

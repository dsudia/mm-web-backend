package algorithm

import "testing"

// TestSomeMatchSuccess tests if someMatch works when there should be a match
func TestSomeMatchSuccess(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 4, 5}
	b := someMatch(s1, s2)
	if b != true {
		t.Error("For", "boolean value",
			"expected", "true",
			"got", b)
	}
}

// TestSomeMatchFailure tests if someMatch works when there should not be a match
func TestSomeMatchFailure(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	b := someMatch(s1, s2)
	if b != false {
		t.Error("For", "boolean value",
			"expected", "false",
			"got", b)
	}
}

// TestCountNumOfMatches should return correct number of matches when there are no matches
func TestCountNumOfMatchesZero(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	n := CountNumOfMatches(s1, s2)
	if n != 0 {
		t.Error("For", "number of matches",
			"expected", "0",
			"got", n)
	}
}

// TestCountNumOfMatches should return correct number of matches when there are matches
func TestCountNumOfMatchesTwo(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 6}
	n := CountNumOfMatches(s1, s2)
	if n != 2 {
		t.Error("For", "number of matches",
			"expected", "0",
			"got", n)
	}
}

// TestFindDecimal checks if findDecimal returns the correct decimal
func TestFindDecimal(t *testing.T) {
	d := findDecimal(3, 5)
	if d != 0.6 {
		t.Error("For", "decimal",
			"expected", "0.6",
			"got", d)
	}
}

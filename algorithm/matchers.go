package algorithm

// MatchAge checks match percent two age slices
func matchAge(memOneAge, memTwoAge []int) float64 {
	count := CountNumOfMatches(memOneAge, memTwoAge)
	if count != 0 {
		return findDecimal(float64(count), float64(len(memOneAge)))
	}
	return float64(-1)
}

// MatchCal checks match percent of two cal slices
func matchCal(memOneCal, memTwoCal []int) float64 {
	if someMatch(memOneCal, memTwoCal) == true {
		return float64(1)
	}
	return float64(0)
}

// MatchEd checks match percent of two ed slices
func matchEd(memOneEd, memTwoEd []int) float64 {
	if someMatch(memOneEd, memTwoEd) == true {
		return float64(1)
	}
	return float64(0)
}

// MatchLoc returns match percent of two loc slices
func matchLoc(memOneLoc, memTwoLoc []int) float64 {
	if someMatch(memOneLoc, memTwoLoc) == true {
		return float64(1)
	}
	return float64(0)
}

// MatchOrg returns match percent of two org slices
func matchOrg(memOneOrg, memTwoOrg []int) float64 {
	if someMatch(memOneOrg, memTwoOrg) == true {
		return float64(1)
	}
	return float64(0)
}

// MatchSize returns match percent of two size slices
func matchSize(memOneSize, memTwoSize) {
	if someMatch(memOneSize, memTwoSize) == true {
		return float64(1)
	}
	return float64(0)
}

// MatchState returns match percent of two state slices
func matchState(memOneState, memTwoState) {
	if someMatch(memOneState, memTwoState) == true {
		return float64(1)
	}
	return float64(0)
}

// MatchTraining returns match percent of two training slices
func matchTraining(memOneTr, memTwoTr) {
	if someMatch(memOneTr, memTwoTr) == true {
		return float64(1)
	}
	return float64(-1)
}

// MatchTraits checks match percent two traits slices
func matchTraits(memOneTraits, memTwoTraits []int) float64 {
	count := CountNumOfMatches(memOneTraits, memTwoTraits)
	if count != 0 {
		return findDecimal(float64(count), float64(len(memOneTraits)))
	}
	return float64(0)
}

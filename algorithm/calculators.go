package algorithm

import "math"

// someMatch checks if any value in one slice exists in another slice
func someMatch(memOneSl, memTwoSl []int) bool {
	for i := range memTwoSl {
		for j := range memOneSl {
			if memTwoSl[i] == memOneSl[j] {
				return true
			}
		}
	}
	return false
}

// CountNumOfMatches checks number of elements that exist in both of two slices
func countNumOfMatches(memOneSl, memTwoSl []int) int {
	num := 0
	for i := range memTwoSl {
		for j := range memOneSl {
			if memTwoSl[i] == memOneSl[j] {
				num++
			}
		}
	}
	return num
}

// findDecimal gets a decimal of one number divided by another
func findDecimal(divisor float64, dividend float64) float64 {
	f := float64(divisor) / float64(dividend)
	return f
}

// MatchPercentOneWay finds match percent for one user against another
func matchPercentOneWay(age, cal, loc, org, size, state, training, traits, ed float64, ageWgt, calWgt, locWgt, orgWgt, sizeWgt, stateWgt, trainingWgt, traitsWgt, edWgt int) {
	ageScore := float64(ageWgt) * age
	calScore := float64(calWgt) * cal
	locScore := float64(locWgt) * loc
	orgScore := float64(orgWgt) * org
	sizeScore := float64(sizeWgt) * size
	stateScore := float64(stateWgt) * state
	trainingScore := float64(trainingWgt) * training
	traitsScore := float64(traitsWgt) * traits
	edScore := float64(edWgt) * ed

	score := ageScore + calScore + locScore + sizeScore + stateScore + trainingScore + traitsScore + edScore

	divisor := ageWgt + calWgt + locWgt + orgWgt + sizeWgt + trainingWgt + traitsWgt + edWgt

	return findDecimal(score, float64(divisor))
}

// MatchPercentMutual finds square root of product of two one way match scores
func matchPercentMutual(score1, score2 float64) float64 {
	return math.Sqrt(score1 * score2)
}

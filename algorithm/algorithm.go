package algorithm

import (
	"math"

	"github.com/dsudia/mm-web-backend/db/models"
)

// Match returns two way matche precentage for two members
func match(ed *models.Educator, school *models.SchoolMatchingProfile) float64 {
	ageMatch := matchAge(school.AgeRanges, ed.AgeRanges)
	if ageMatch == -1 {
		return 0
	}

	stateMatch := matchState(school.States, ed.States)
	if stateMatch == -1 {
		return 0
	}

	trainingMatch := matchTraining(ed.Trainings, school.Trainings)
	if trainingMatch == -1 {
		return 0
	}

	traitMatch := matchTraits(ed.Traits, school.Traits)
	calMatch := matchCal(ed.Cals, school.Cals)
	locMatch := matchLocs(ed.LocTypes, school.LocTypes)
	orgMatch := matchOrg(ed.OrgTypes, school.OrgTypes)
	sizeMatch := matchSize(ed.Sizes, school.Sizes)
	edMatch := matchEd(ed.EdTypes, school.EdTypes)

	matchPercentEd := matchPercentOneWay(ageMatch, calMatch, locMatch, orgMatch, sizeMatch, stateMatch, trainingMatch, traitsMatch, edMatch, ed.AgesRangesWgt, ed.CalsWgt, ed.LocTypesWgt, ed.OrgTypesWgt, ed.SizesWgt, ed.StatesWgt, ed.TrainingsWgt, ed.TraitsWgt, ed.EdTypesWgt)

	matchPercentSchool := matchPercentOneWay(ageMatch, calMatch, locMatch, orgMatch, sizeMatch, stateMatch, trainingMatch, traitsMatch, edMatch, school.AgesRangesWgt, school.CalsWgt, school.LocTypesWgt, school.OrgTypesWgt, school.SizesWgt, school.StatesWgt, school.TrainingsWgt, school.TraitsWgt, school.EdTypesWgt)

	matchPercent := matchPercentMutual(matchPercentEd, matchPercentSchool)

	shift := math.Pow(10, float64(2))
	num := (matchPercent * shift)
	toTwoPlaces := math.Floor(num+.5) / shift

	return toTwoPlaces
}

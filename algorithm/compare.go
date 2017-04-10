package algorithm

import (
	"github.com/dsudia/mm-web-backend/db"
	"github.com/dsudia/mm-web-backend/models"
)

// RunMatchComparisonEducator runs match comparisons and updates db for educator
func RunMatchComparisonEducator(id int) {
	matchPercent := 0
	ed := db.GetEducator(id)
	smps := db.GetSchoolMatchingProfiles()

	for smp := range smps {
		existing := db.GetMatchesByEducatorAndSchool(ed.ID, smp.ID)
		if len(existing) == 0 {
			percent := match(ed, smp) * 100
			if percent > 80 {
				m := models.Match{
					SchoolMatchingProfileID: smp.ID,
					EducatorID:              ed.ID,
					Percentage:              percent,
					EducatorConfirmation:    false,
					SchoolConfirmation:      false,
					EducatorDenial:          false,
					SchoolDenial:            false,
				}
				mid, _ := db.CreateMatch(m)
			}
		}
	}
	return
}

// RunMatchComparisonSchool runs match comparisons and updates db for school
func RunMatchComparisonSchool(id int) {
	matchPercent := 0
	smp := db.GetSchoolMatchingProfile(id)
	eds := db.GetEducators()

	for ed := range eds {
		existing := db.GetMatchesByEducatorAndSchool(ed.ID, smp.ID)
		if len(existing) == 0 {
			percent := match(ed, smp) * 100
			if percent > 80 {
				m := models.Match{
					SchoolMatchingProfileID: smp.ID,
					EducatorID:              ed.ID,
					Percentage:              percent,
					EducatorConfirmation:    false,
					SchoolConfirmation:      false,
					EducatorDenial:          false,
					SchoolDenial:            false,
				}
				mid, _ := db.CreateMatch(m)
			}
		}
	}

}

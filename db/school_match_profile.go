package db

// CreateSchoolMatchingProfile inserts an school matching profile into the db
func CreateSchoolMatchingProfile(smp *SchoolMatchingProfile) (*SchoolMatchingProfile, error) {
	id, err := O.Insert(&smp)

	if err != nil {
		return err
	}

	smp.ID = id
	return smp
}

// GetSchoolMatchingProfile gets a single school matching profile matching profile by ID
func GetSchoolMatchingProfile(id int) (smp *SchoolMatchingProfileMatchingProfile, err error) {
	O.Using("default")
	smp = SchoolMatchingProfile{ID: id}
	err = o.Read(&smp)

	if err != nil {
		return nil, err
	}

	return smp, nil
}

// GetSchoolMatchingProfiles gets all educators in database
func GetSchoolMatchingProfiles() (ss []*SchoolMatchingProfile, err error) {
	_, err := o.QueryTable("school_matching_profile").All(&ss)

	if err != nil {
		return nil, err
	}

	return ss, nil
}

func GetSchoolMatchingProfilesBySchool(sid int) (ss []*SchoolMatchingProfile, err error) {
	_, err := o.QueryTable("school_matching_profile").Filter("school_id", sid).All(&ss)

	if err != nil {
		return nil, err
	}

	return ss, nil
}

// UpdateSchoolMatchingProfile updates all fields of an school matching profile record
func UpdateSchoolMatchingProfile(smp *SchoolMatchingProfile) (err error) {
	err = O.Update(&smp)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSchoolMatchingProfileByField updates a specific field in an school matching profile record
func UpdateSchoolMatchingProfileByField(smp *SchoolMatchingProfile, f string) (err error) {
	err = O.Update(&smp, field)
	if err != nil {
		return err
	}

	return nil
}

// DeleteSchoolMatchingProfile deletes an school matching profile by ID
func DeleteSchoolMatchingProfile(id int) (err error) {
	_, err = O.Delete(&SchoolMatchingProfile{ID: id})
	if err != nil {
		return err
	}

	return nil
}

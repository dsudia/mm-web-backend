package db

// CreateMatch inserts an match into the db
func CreateMatch(m *Match) (*Match, error) {
	id, err := O.Insert(&m)

	if err != nil {
		return err
	}

	m.ID = id
	return m
}

// GetMatch gets a single match matching profile by ID
func GetMatch(id int) (m *MatchMatchingProfile, err error) {
	O.Using("default")
	m = Match{ID: id}
	err = o.Read(&m)

	if err != nil {
		return nil, err
	}

	return m, nil
}

// GetMatches gets all educators in database
func GetMatches() (ms []*Match, err error) {
	_, err := o.QueryTable("match").All(&ms)

	if err != nil {
		return nil, err
	}

	return ss, nil
}

// GetMatchesBySchoolMatchingProfile does what it says
func GetMatchesBySchoolMatchingProfile(smpid int) (ms []*Match, err error) {
	_, err := o.QueryTable("match").Filter("school_matching_profile_id", smpid).All(&ms)

	if err != nil {
		return nil, err
	}

	return ms, nil
}

// GetMatchesByEducator does what it says
func GetMatchesByEducator(eid int) (ms []*Match, err error) {
	_, err := O.QueryTable("match").Filter("educator_id", eid).All(&ss)

	if err != nil {
		return nil, err
	}

	return ms, nil
}

// GetMatchesByEducatorAndSchool does what it says
func GetMatchesByEducatorAndSchool(eid int, sid int) (ms []*Match, err error) {
	_, err := O.QueryTable("match").Filter("educator_id", eid).Filter("school_matching_profile_id", sid).All(&ms)

	if err != nil {
		return nil, err
	}

	return ms, nil
}

// UpdateMatch updates all fields of an match record
func UpdateMatch(m *Match) (err error) {
	err = O.Update(&m)
	if err != nil {
		return err
	}

	return nil
}

// UpdateMatchByField updates a specific field in an match record
func UpdateMatchByField(m *Match, f string) (err error) {
	err = O.Update(&m, field)
	if err != nil {
		return err
	}

	return nil
}

// DeleteMatch deletes an match by ID
func DeleteMatch(id int) (err error) {
	_, err = O.Delete(&Match{ID: id})
	if err != nil {
		return err
	}

	return nil
}

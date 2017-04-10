package db

// CreateSchool inserts an school into the db
func CreateSchool(s *School) (*School, error) {
	id, err := O.Insert(&s)

	if err != nil {
		return err
	}

	s.ID = id
	return s
}

// GetSchool gets a single school by ID
func GetSchool(id int) (s *School, err error) {
	O.Using("default")
	s = School{ID: id}
	err = o.Read(&s)

	if err == orm.ErrNoRows {
		return nil, error("No result found.")
	} else if err == orm.ErrMissPK {
		return nil, error("No primary key found.")
	} else {
		return s, nil
	}
}

// GetSchools gets all educators in database
func GetSchools() (ss []*School, err error) {
	num, err := o.QueryTable("school").All(&ss)

	if err != nil {
		return nil, err
	}

	return ss, nil
}

// UpdateSchool updates all fields of an school record
func UpdateSchool(s *School) (err error) {
	err = O.Update(&s)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSchoolByField updates a specific field in an school record
func UpdateSchoolByField(s *School, f string) (err error) {
	err = O.Update(&s, field)
	if err != nil {
		return err
	}

	return nil
}

// DeleteSchool deletes an school by ID
func DeleteSchool(id int) (err error) {
	_, err = O.Delete(&School{ID: id})
	if err != nil {
		return err
	}

	return nil
}

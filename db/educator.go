package db

// CreateEducator inserts an educator into the db
func CreateEducator(e *Educator) (*Educator, err error) {
	id, err = O.Insert(&e)

	if err != nil {
		return err
	}

	e.ID = id
	return e
}

// GetEducator gets a single educator by ID
func GetEducator(id int) (e *Educator, err error) {
	O.Using("default")
	e = Educator{ID: id}
	err = o.Read(&e)

	if err == orm.ErrNoRows {
		return nil, error("No result found.")
	} else if err == orm.ErrMissPK {
		return nil, error("No primary key found.")
	} else {
		return e, nil
	}
}

// GetEducators gets all educators in database
func GetEducators() (es []*Educator, err error) {
	num, err := o.QueryTable("educator").All(&es)

	if err != nil {
		return nil, err
	}

	return es, nil
}

// UpdateEducator updates all fields of an educator record
func UpdateEducator(e *Educator) (err error) {
	err = O.Update(&e)
	if err != nil {
		return err
	}

	return nil
}

// UpdateEducatorByField updates a specific field in an educator record
func UpdateEducatorByField(e *Educator, f string) (err error) {
	err = O.Update(&e, field)
	if err != nil {
		return err
	}

	return nil
}

// DeleteEducator deletes an educator by ID
func DeleteEducator(id int) (err error) {
	_, err = O.Delete(&Educator{ID: id})
	if err != nil {
		return err
	}

	return nil
}

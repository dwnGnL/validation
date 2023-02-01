package repository

func (r RepoImpl) GetTestTable() (*TestTable, error) {
	testTable := new(TestTable)
	err := r.db.Find(testTable).Error
	if err != nil {
		return nil, err
	}
	return testTable, nil
}

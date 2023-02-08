package repository

import (
	"errors"
	"log"
)

func (r RepoImpl) GetTestTable() (*TestTable, error) {
	testTable := new(TestTable)
	err := r.db.Find(testTable).Error
	if err != nil {
		return nil, err
	}
	return testTable, nil
}

func (r RepoImpl) Registration(info *Users) error {
	tx := r.db.Create(info)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r RepoImpl) Login(info string) (*Users, error) {
	var user Users

	// ============ todo ask question " why isnt working?"
	//tx := r.db.Table("users").Scan(&user).Where(info)
	//if tx.Error != nil {
	//	//log.Println(error)
	//	return nil, errors.New("tx.error")
	//}
	sqlQuery := `select name, password, id from users where login = ?;`

	if err := r.db.Raw(sqlQuery, info).Scan(&user).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	if user.Name == "" {

		return nil, errors.New("test")

	}

	return &user, nil
}

func (r RepoImpl) SaveToken(token Tokens) error {

	tx := r.db.Create(&token)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}

	return nil
}

//func (r RepoImpl) AddIntoTable() (*TestTable, error) {
//	//tx := r.db.Create(info)
//	//if tx.Error != nil {
//	//	log.Println(tx.Error)
//	//	return tx.Error
//	//}
//	//return nil
//	testTable := new(TestTable)
//	err := r.db.Find(testTable).Error
//	if err != nil {
//		return nil, err
//	}
//	return testTable, nil
//}

package repository

type TestTable struct {
	ID    int64  `gorm:"column:id;primary_key;autoIncrement"`
	Title string `gorm:"column:title"`
}

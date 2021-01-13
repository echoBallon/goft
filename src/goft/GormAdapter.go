
package goft

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type GormAdapter struct {
	*gorm.DB
}
func NewGormAdapter() *GormAdapter {
	db, err := gorm.Open("mysql",
		"root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	Error(err)
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)
	return &GormAdapter{DB:db}
}
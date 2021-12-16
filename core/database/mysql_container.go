package database

import (
	"gorm.io/driver/mysql"

	"gorm.io/gorm"

	"github.com/renjingneng/goapp/core/config"
)

var mysqlContainer map[string]*gorm.DB

//GetEntityFromMysqlContainer is
func GetEntityFromMysqlContainer(database string, mode string) *gorm.DB {
	if database == "" || mode == "" {
		return nil
	}
	dbname := database + mode
	if db, ok := mysqlContainer[dbname]; ok {
		return db
	}
	if ok := config.Get(dbname); ok == "" {
		return nil
	}
	if db, err := gorm.Open(mysql.Open(config.Get(dbname)), &gorm.Config{}); err != nil {
		return nil
	} else {
		mysqlContainer[dbname] = db
		return db
	}
}
func init() {
	if mysqlContainer == nil {
		mysqlContainer = make(map[string]*gorm.DB)
	}
}

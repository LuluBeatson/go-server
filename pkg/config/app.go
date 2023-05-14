package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/LuluBeatson/go-server/pkg/utils"
)

var (
	db *gorm.DB
)

func Connect() {
	config, err := utils.GetConfig()
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/bookstore?charset=utf8&parseTime=True&loc=Local", config.MySQL.User, config.MySQL.Password, config.MySQL.Host, config.MySQL.Port)
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

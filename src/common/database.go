package common

import (
	"fmt"

	"../configuration"
	"github.com/jinzhu/gorm"
)

//DB is the pointer of database instance and applies the rules
var DB *gorm.DB

//DbInit opens db connection
func DbInit() *gorm.DB {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		switch defaultTableName {
		case "liquidances":
			return "public.liquidance"
		default:
			return "public." + defaultTableName
		}

	}

	db, err := gorm.Open("postgres", configuration.GetDbConnectionString())
	if err != nil {
		fmt.Println("Can't connect to db")
	}

	DB = db
	return DB
}

//GetDbInstance returns the pointer of db
func GetDbInstance() *gorm.DB {
	return DB
}

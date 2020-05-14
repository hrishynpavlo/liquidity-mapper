package common

import (
	"fmt"

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

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=Ph@!436623 sslmode=disable")
	if err != nil {
		fmt.Println("Can't connect to db")
	}

	DB = db
	return DB
}

//GetInstance returns the pointer of db
func GetInstance() *gorm.DB {
	return DB
}

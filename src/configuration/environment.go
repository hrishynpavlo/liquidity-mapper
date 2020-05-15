package configuration

import (
	"fmt"

	"github.com/joho/godotenv"
)

//ENV structure of configuration
type ENV struct {
	Db *DbEnv
}

//DbEnv structure of database configuration
type DbEnv struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

var env *ENV

//SetUp reads configuration from file and fill structure
func SetUp() {
	mapEnv, error := godotenv.Read()
	if error != nil {
		panic("Configuration is lost")
	}
	var dbEnv = DbEnv{Port: mapEnv["DB_PORT"], Host: mapEnv["DB_HOST"], Database: mapEnv["DB_NAME"], User: mapEnv["DB_USER"], Password: mapEnv["DB_PASSWORD"]}
	var envP = ENV{Db: &dbEnv}
	env = &envP
}

//GetDbConnectionString returns connection string acording to configuration
func GetDbConnectionString() string {
	return fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", env.Db.Host, env.Db.Port, env.Db.User, env.Db.Database, env.Db.Password)
}

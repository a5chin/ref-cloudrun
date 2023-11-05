package driver

import (
	"database/sql"
	"fmt"
	"os"
	"ref/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const TxKey = "transactionObject"
const ErrDuplicateEntryNumber = 1062

func NewDB(conf *config.Config) *gorm.DB {
	dbUser := conf.DB_USER
	dbPwd := conf.DB_PWD
	dbName := conf.DB_NAME
	dbTCPHost := conf.DB_TCPHOST
	dbPort := conf.DB_PORT
	dbURI := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPwd, dbTCPHost, dbPort, dbName,
	)

	sqlDB, err := sql.Open("mysql", dbURI)
	dbPool, err := gorm.Open(
		mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return dbPool
}

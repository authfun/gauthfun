package database

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/authfun/gauthfun/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var AuthDatabase *gorm.DB

func init() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		config.AppConfig.MySql.User,
		config.AppConfig.MySql.Password,
		config.AppConfig.MySql.Host,
		config.AppConfig.MySql.Port,
		config.AppConfig.MySql.DbName,
		config.AppConfig.MySql.Parameters,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("fatal error open mysql connection of auth database: %s", err))
	}

	initDb(db)
	AuthDatabase = db
}

func initDb(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadFile("./data/init.sql")
	if err != nil {
		panic(err)
	}

	scripts := strings.Split(string(bytes), ";")
	for _, script := range scripts {
		if len(script) <= 0 {
			continue
		}
		_, err := sqlDB.Exec(script)
		if err != nil {
			panic(err)
		}
	}
}

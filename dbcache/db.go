package dbcache

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func initDB(url string, debug bool, maxIdleConn, maxOpenConn int) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	db.LogMode(debug)

	db.DB().SetConnMaxLifetime(2 * time.Hour)
	if maxIdleConn != 0 {
		db.DB().SetMaxIdleConns(maxIdleConn)
	}
	if maxOpenConn != 0 {
		db.DB().SetMaxOpenConns(maxOpenConn)
	}
	if err := db.DB().Ping(); err != nil {
		return db, err
	}
	return db, nil
}

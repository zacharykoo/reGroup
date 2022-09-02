// package database will provide a database connection to sqlite for now
package database

import (
	"github.com/glebarez/sqlite"
	"github.com/zacharykoo/reGroup/backend/pkg/model"
	"gorm.io/gorm"
)

func ConnectSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("reGroup.sqlite"), &gorm.Config{})
	return db, err
}

func MigrateTables(db *gorm.DB) error {
	err := db.AutoMigrate(
		model.User{},
	)
	return err
}

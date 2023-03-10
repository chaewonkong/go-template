package config

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteDialector gorm.Dialector

func NewSQLiteDialector(settings *Settings) SQLiteDialector {
	db := fmt.Sprintf("%s.db", settings.Database.Name)
	return sqlite.Open(db)
}

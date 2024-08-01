package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sqlite struct {
	*gorm.DB
}

func New(database_path string) (*Sqlite, error) {
	db, err := gorm.Open(sqlite.Open(database_path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Sqlite{db}, nil
}

package app

import (
	"fmt"

	"github.com/vicdevcode/exam/internal/entity"
	"github.com/vicdevcode/exam/internal/sqlite"
)

func Migrate(runType string, db *sqlite.Sqlite) {
	switch runType {
	case "create":
		if err := create(db); err != nil {
			panic(err)
		}
	case "drop":
		if err := drop(db); err != nil {
			panic(err)
		}
	case "reset":
		if err := drop(db); err != nil {
			panic(err)
		}
		if err := create(db); err != nil {
			panic(err)
		}
	default:
		panic("?")
	}
}

func create(db *sqlite.Sqlite) error {
	if err := db.AutoMigrate(
		&entity.SubCategory{},
		&entity.Card{},
		&entity.Category{},
	); err != nil {
		return err
	}
	return nil
}

func drop(db *sqlite.Sqlite) error {
	tables := []string{"items", "categories", "sub_categories"}
	for _, t := range tables {
		if err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", t)).Error; err != nil {
			return err
		}
	}
	return nil
}

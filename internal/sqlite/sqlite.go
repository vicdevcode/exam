package sqlite

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/vicdevcode/exam/internal/app/config"
)

type Sqlite struct {
	*gorm.DB
}

func New(cfg *config.Postgres) (*Sqlite, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Sqlite{db}, nil
}

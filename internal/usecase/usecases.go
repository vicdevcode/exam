package usecase

import (
	"github.com/vicdevcode/exam/internal/app/config"
	"github.com/vicdevcode/exam/internal/sqlite"
)

type UseCases struct {
	CardUseCase        Card
	CategoryUseCase    Category
	SubCategoryUseCase SubCategory
}

func New(cfg *config.Config, db *sqlite.Sqlite) UseCases {
	t := cfg.ContextTimeout
	return UseCases{
		CardUseCase:        NewCard(db, t),
		CategoryUseCase:    NewCategory(db, t),
		SubCategoryUseCase: NewSubCategory(db, t),
	}
}

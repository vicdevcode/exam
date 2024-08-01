package usecase

import (
	"context"

	"github.com/vicdevcode/exam/internal/entity"
)

type (
	Card interface {
		Create(context.Context, entity.Card) (*entity.Card, error)
		Update(context.Context, entity.Card) (*entity.Card, error)
		Delete(context.Context, uint) error
		FindAllBySubCategoryID(context.Context, uint) ([]entity.Card, error)
	}
	Category interface {
		Create(context.Context, entity.Category) (*entity.Category, error)
		Update(context.Context, entity.Category) (*entity.Category, error)
		Delete(context.Context, uint) error
		FindAll(context.Context) ([]entity.Category, error)
	}
	SubCategory interface {
		Create(context.Context, entity.SubCategory) (*entity.SubCategory, error)
		Update(context.Context, entity.SubCategory) (*entity.SubCategory, error)
		Delete(context.Context, uint) error
		FindAllByCategoryID(context.Context, uint) ([]entity.SubCategory, error)
	}
)

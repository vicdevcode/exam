package usecase

import (
	"time"

	"golang.org/x/net/context"

	"github.com/vicdevcode/exam/internal/entity"
	"github.com/vicdevcode/exam/internal/sqlite"
)

type CategoryUseCase struct {
	db         *sqlite.Sqlite
	ctxTimeout time.Duration
}

func NewCategory(db *sqlite.Sqlite, t time.Duration) *CategoryUseCase {
	return &CategoryUseCase{db, t}
}

func (uc *CategoryUseCase) Create(
	c context.Context,
	body entity.Category,
) (*entity.Category, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	category := &entity.Category{
		Name: body.Name,
	}

	if err := uc.db.WithContext(ctx).Create(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (uc *CategoryUseCase) FindAll(
	c context.Context,
) ([]entity.Category, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	var categories []entity.Category

	if err := uc.db.WithContext(ctx).Preload("SubCategories.Cards").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (uc *CategoryUseCase) Update(
	c context.Context, body entity.Category,
) (*entity.Category, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	category := &entity.Category{
		ID: body.ID,
	}

	if err := uc.db.WithContext(ctx).Model(category).Updates(entity.Category{Name: body.Name}).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (uc *CategoryUseCase) Delete(ctx context.Context, id uint) error {
	return uc.db.WithContext(ctx).Unscoped().Delete(&entity.Category{ID: id}).Error
}

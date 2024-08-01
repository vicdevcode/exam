package usecase

import (
	"time"

	"golang.org/x/net/context"

	"github.com/vicdevcode/exam/internal/entity"
	"github.com/vicdevcode/exam/internal/sqlite"
)

type SubCategoryUseCase struct {
	db         *sqlite.Sqlite
	ctxTimeout time.Duration
}

func NewSubCategory(db *sqlite.Sqlite, t time.Duration) *SubCategoryUseCase {
	return &SubCategoryUseCase{db, t}
}

func (uc *SubCategoryUseCase) Create(
	c context.Context,
	body entity.SubCategory,
) (*entity.SubCategory, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	category := &entity.SubCategory{
		Name:       body.Name,
		CategoryID: body.CategoryID,
	}

	if err := uc.db.WithContext(ctx).Create(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (uc *SubCategoryUseCase) FindAllByCategoryID(
	c context.Context,
	id uint,
) ([]entity.SubCategory, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	var categories []entity.SubCategory

	if err := uc.db.WithContext(ctx).Where("category_id = ?", id).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (uc *SubCategoryUseCase) Update(
	c context.Context, body entity.SubCategory,
) (*entity.SubCategory, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	subCategory := &entity.SubCategory{
		ID: body.ID,
	}

	if err := uc.db.WithContext(ctx).Model(subCategory).Updates(entity.SubCategory{Name: body.Name, CategoryID: body.CategoryID}).Error; err != nil {
		return nil, err
	}
	return subCategory, nil
}

func (uc *SubCategoryUseCase) Delete(ctx context.Context, id uint) error {
	return uc.db.WithContext(ctx).Unscoped().Delete(&entity.SubCategory{ID: id}).Error
}

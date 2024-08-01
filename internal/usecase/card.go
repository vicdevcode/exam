package usecase

import (
	"time"

	"golang.org/x/net/context"

	"github.com/vicdevcode/exam/internal/entity"
	"github.com/vicdevcode/exam/internal/sqlite"
)

type CardUseCase struct {
	db         *sqlite.Sqlite
	ctxTimeout time.Duration
}

func NewCard(db *sqlite.Sqlite, t time.Duration) *CardUseCase {
	return &CardUseCase{db, t}
}

func (uc *CardUseCase) Create(c context.Context, body entity.Card) (*entity.Card, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	card := &entity.Card{
		Question:      body.Question,
		Answer:        body.Answer,
		SubCategoryID: body.SubCategoryID,
	}

	if err := uc.db.WithContext(ctx).Create(&card).Error; err != nil {
		return nil, err
	}
	return card, nil
}

func (uc *CardUseCase) FindAllBySubCategoryID(c context.Context, id uint) ([]entity.Card, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	var cards []entity.Card

	if err := uc.db.WithContext(ctx).Where("sub_category_id = ?", id).Find(&cards).Error; err != nil {
		return nil, err
	}
	return cards, nil
}

func (uc *CardUseCase) Update(
	c context.Context, body entity.Card,
) (*entity.Card, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	category := &entity.Card{
		ID: body.ID,
	}

	if err := uc.db.WithContext(ctx).Model(category).Updates(entity.Card{Question: body.Question, Answer: body.Answer, SubCategoryID: body.SubCategoryID}).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (uc *CardUseCase) Delete(ctx context.Context, id uint) error {
	return uc.db.WithContext(ctx).Unscoped().Delete(&entity.Card{ID: id}).Error
}

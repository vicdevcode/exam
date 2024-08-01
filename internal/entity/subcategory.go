package entity

type SubCategory struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Cards      []Card `json:"cards"`
	CategoryID uint   `json:"category_id"`
}

package entity

type Card struct {
	ID            uint   `json:"id"`
	Question      string `json:"question"`
	Answer        string `json:"answer"`
	SubCategoryID uint   `json:"sub_category_id"`
}

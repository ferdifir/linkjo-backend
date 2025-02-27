package models

type SubCategory struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID            uint          `json:"id"`
	Name          string        `json:"name"`
	SubCategories []SubCategory `json:"subcategories"`
}

package models

type Product struct {
	ID          string  `gorm:"primaryKey;column:id"`
	Title       string  `gorm:"column:title"`
	Price       float64 `gorm:"column:price"`
	Description string  `gorm:"column:description"`
	Category    string  `gorm:"column:category"`
	Image       string  `gorm:"column:image"`
}

type DataTemp struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
}

package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID         uint    `json:"id" gorm:"primary_key"`
	Name       string  `json:"name"`
	CategoryID uint    `json:"category_id"`
	Price      float64 `json:"price"`
}

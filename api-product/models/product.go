package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:varchar(255)"`
}

type ProductResponse struct {
	ID          int    `json:"id" gorm:"-all"`
	Name        string `json:"name" gorm:"-all"`
	Description string `json:"description" gorm:"-all"`
}

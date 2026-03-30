package product

import (
	"gorm.io/gorm"

	"github.com/lib/pq"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
}

func NewProduct(name string, description string) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Images:      pq.StringArray{},
	}
}

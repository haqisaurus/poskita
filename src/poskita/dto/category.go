package dto

import "time"

type ProductCategoryRq struct {
	ID          uint64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"` 
}

type ProductCategoryRs struct {
	ID          uint64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedBy   string    `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
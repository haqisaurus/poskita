package dto

import "time"

type ProductSupplierRq struct {
	ID          uint64     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Phone       string     `json:"phone"`
	Address     string     `json:"address"`
}

type ProductSupplierRs struct {
	ID          uint64     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Phone       string     `json:"phone"`
	Address     string     `json:"address"`
	CreatedBy   string     `json:"created_by"`
	CreatedAt   time.Time  ` json:"created_at"`
	UpdatedBy   string     `json:"updated_by"`
	UpdatedAt   time.Time  ` json:"updated_at,omitempty"`
}
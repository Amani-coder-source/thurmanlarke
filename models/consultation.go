package models

import "time"

type Consultation struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Service     string    `json:"service"`
	BusinessName string   `json:"business_name"`
	BusinessType string   `json:"business_type"`
	Description string    `json:"description"`
	Budget      string    `json:"budget"`
	PreferredDate string  `json:"preferred_date"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
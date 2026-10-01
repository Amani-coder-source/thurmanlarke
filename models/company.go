package models

import "time"

type Company struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	CompanyName      string    `json:"company_name"`
	CompanyType      string    `json:"company_type"`
	BusinessActivity string    `json:"business_activity"`
	Address          string    `json:"address"`
	City             string    `json:"city"`
	Country          string    `json:"country"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
}
package repository

import (
	"database/sql"

	"Thurmanlark/models"
)

type ConsultationRepository struct {
	DB *sql.DB
}

// CreateConsultation saves a new consultation request.
func (r ConsultationRepository) CreateConsultation(
	consultation models.Consultation,
) error {

	query := `
		INSERT INTO consultations (
			user_id,
			service,
			business_name,
			business_type,
			description,
			budget,
			preferred_date,
			status
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.DB.Exec(
		query,
		consultation.UserID,
		consultation.Service,
		consultation.BusinessName,
		consultation.BusinessType,
		consultation.Description,
		consultation.Budget,
		consultation.PreferredDate,
		consultation.Status,
	)

	return err
}
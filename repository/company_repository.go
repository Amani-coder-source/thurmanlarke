package repository

import (
	"database/sql"

	"Thurmanlark/models"
)

type CompanyRepository struct {
	DB *sql.DB
}

// CreateCompany saves a company registration application.
func (r CompanyRepository) CreateCompany(company models.Company) error {

	query := `
		INSERT INTO companies (
			user_id,
			company_name,
			company_type,
			business_activity,
			address,
			city,
			country,
			status
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.DB.Exec(
		query,
		company.UserID,
		company.CompanyName,
		company.CompanyType,
		company.BusinessActivity,
		company.Address,
		company.City,
		company.Country,
		company.Status,
	)

	return err
}

// GetCompaniesByUserID gets all company applications
// belonging to a particular user.
func (r CompanyRepository) GetCompaniesByUserID(userID int) ([]models.Company, error) {

	query := `
		SELECT
			id,
			user_id,
			company_name,
			company_type,
			business_activity,
			address,
			city,
			country,
			status,
			created_at
		FROM companies
		WHERE user_id = ?
		ORDER BY id DESC
	`

	rows, err := r.DB.Query(query, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var companies []models.Company

	for rows.Next() {

		var company models.Company

		err := rows.Scan(
			&company.ID,
			&company.UserID,
			&company.CompanyName,
			&company.CompanyType,
			&company.BusinessActivity,
			&company.Address,
			&company.City,
			&company.Country,
			&company.Status,
			&company.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}
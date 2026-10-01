package repository

import (
	"database/sql"
	"time"

	"Thurmanlark/models"
)

type UserRepository struct {
	DB *sql.DB
}

// CreateUser saves a new user in the database.
func (r UserRepository) CreateUser(user models.User) error {

	query := `
		INSERT INTO users (
			name,
			email,
			password,
			role
		)
		VALUES (?, ?, ?, ?)
	`

	_, err := r.DB.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
	)

	return err
}

// GetUserByEmail finds a user by email.
func (r UserRepository) GetUserByEmail(email string) (models.User, error) {

	var user models.User

	var createdAt []byte

	query := `
		SELECT
			id,
			name,
			email,
			password,
			role,
			created_at
		FROM users
		WHERE email = ?
	`

	err := r.DB.QueryRow(
		query,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&createdAt,
	)

	if err != nil {
		return user, err
	}

	// Convert MySQL created_at value to time.Time.
	if len(createdAt) > 0 {
		parsedTime, err := time.Parse(
			"2006-01-02 15:04:05",
			string(createdAt),
		)

		if err != nil {
			return user, err
		}

		user.CreatedAt = parsedTime
	}

	return user, nil
}// GetUserByID finds a user by their ID.
func (r UserRepository) GetUserByID(id int) (models.User, error) {

	var user models.User

	var createdAt []byte

	query := `
		SELECT
			id,
			name,
			email,
			password,
			role,
			created_at
		FROM users
		WHERE id = ?
	`

	err := r.DB.QueryRow(
		query,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&createdAt,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}
package dbrepo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/salimmia/go-architecture/internal/application"
	"github.com/salimmia/go-architecture/internal/models"
	"github.com/salimmia/go-architecture/internal/repository"
)

type mySqlDbRepo struct{
	App *application.Application
	DB *sql.DB
}

func NewMySqlDbRepo(app *application.Application, db *sql.DB) repository.Database{
	return &mySqlDbRepo{
		App: app,
		DB: db,
	}
}

func (app *mySqlDbRepo) GetAllUsers() ([]models.User, error){
	var users []models.User

	return users, nil
}

func (app *mySqlDbRepo) RegistrationUser(user *models.User) (*models.UserID, error){
	ctx, cancel:= context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	query := `
		INSERT INTO users (email, password, first_name, last_name) VALUES ($1, $2, $3, $4) RETURNING id;
	`
	var userId uuid.UUID
	err := app.DB.QueryRowContext(ctx, query, user.Email, user.Password.Hash, user.FirstName, user.LastName).Scan(&userId)
	if err != nil{
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return nil, models.ErrDuplicateEmail
		default:
			return nil, err
		}
	}

	query = `
		INSERT INTO user_profile (user_id) VALUES ($1) ON CONFLICT (user_id) DO NOTHING RETURNING user_id;
	`
	err = app.DB.QueryRowContext(ctx, query, userId).Scan(&userId)
	if err != nil {
		return nil, err
	}
	id := models.UserID{
		Id: userId,
	}

	return &id, nil
}

func (app *mySqlDbRepo) GetUserById(userId uuid.UUID) (*models.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var user models.User
	var userProfile models.UserProfile

	query_user := `
		SELECT 
			u.*, p.* 
		FROM 
			users u 
			LEFT JOIN user_profile p ON p.user_id = u.id 
		WHERE 
			u.is_active = true AND u.id = $1
	`

	err := app.DB.QueryRowContext(ctx, query_user, userId).Scan(
		&user.ID,
		&user.Email,
		&user.Password.Hash,
		&user.FirstName,
		&user.LastName,
		&user.IsActive,
		&user.IsStaff,
		&user.IsSuperuser,
		&user.Thumbnail,
		&user.CreatedAt,
		&userProfile.ID,
		&userProfile.UserID,
		&userProfile.PhoneNumber,
		&userProfile.BirthDate,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, models.ErrRecordNotFound
		default:
			return nil, errors.New("error")
		}
	}
	user.Profile = userProfile

	return &user, nil
}

func (app *mySqlDbRepo) GetUserByEmail(email string) (*models.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var user models.User
	var userProfile models.UserProfile

	query_user := `
		SELECT 
			u.*, p.* 
		FROM 
			users u 
			LEFT JOIN user_profile p ON p.user_id = u.id 
		WHERE 
			u.is_active = true AND u.email = $1
	`

	err := app.DB.QueryRowContext(ctx, query_user, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password.Hash,
		&user.FirstName,
		&user.LastName,
		&user.IsActive,
		&user.IsStaff,
		&user.IsSuperuser,
		&user.Thumbnail,
		&user.CreatedAt,
		&userProfile.ID,
		&userProfile.UserID,
		&userProfile.PhoneNumber,
		&userProfile.BirthDate,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, models.ErrRecordNotFound
		default:
			return nil, err
		}
	}
	user.Profile = userProfile

	return &user, nil
}

func (app *mySqlDbRepo) UpdateUser(user *models.User) (*models.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var newUser models.User
	var newProfile models.UserProfile

	query_user := `
		UPDATE 
			users 
		SET 
			first_name = COALESCE($1, first_name), 
			last_name = COALESCE($2, last_name), 
			thumbnail = COALESCE($3, thumbnail)
		WHERE 
			id = $4 AND is_active = true
		RETURNING id, email, password, first_name, last_name, is_active, is_staff, is_superuser, thumbnail, created_at;
	`

	err := app.DB.QueryRowContext(ctx, query_user, user.FirstName, user.LastName, user.Thumbnail, user.ID).Scan(
		&newUser.ID, 
		&newUser.Email, 
		&newUser.Password.Hash, 
		&newUser.FirstName,
		&newUser.LastName,
		&newUser.IsActive,
		&newUser.IsStaff,
		&newUser.IsSuperuser,
		&newUser.Thumbnail,
		&newUser.CreatedAt,
	)

	if err != nil{
		return nil, err
	}

	query_user_profile := `
		UPDATE 
			user_profile 
		SET 
			phone_number = NULLIF($1, ''), 
			birth_date = $2::timestamp::date
		WHERE 
			user_id = $3
		RETURNING id, user_id, phone_number, birth_date;
	`
	err = app.DB.QueryRowContext(ctx, query_user_profile, user.Profile.PhoneNumber, user.Profile.BirthDate, user.Profile.UserID).Scan(
		&newProfile.ID, 
		&newProfile.UserID, 
		&newProfile.PhoneNumber, 
		&newProfile.BirthDate,
	)

	if err != nil{
		return nil, err
	}

	newUser.Profile = newProfile
	
	return &newUser, nil
}
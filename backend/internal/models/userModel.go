package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/salimmia/go-architecture/internal/types"
)

type UserProfile struct {
    ID          *uuid.UUID     `json:"id"`
    UserID      *uuid.UUID     `json:"user_id"`
    PhoneNumber *string        `json:"phone_number"`
    BirthDate   types.NullTime `json:"birth_date"`
}

type User struct {
    ID          uuid.UUID       `json:"id"`
    Email       string          `json:"email"`
    Password    Password        `json:"-"`
    FirstName   string          `json:"first_name"`
    LastName    string          `json:"last_name"`
    IsActive    bool            `json:"is_active"`
    IsStaff     bool            `json:"is_staff"`
    IsSuperuser bool            `json:"is_superuser"`
    Thumbnail   *string         `json:"thumbnail"`
    CreatedAt   time.Time       `json:"created_at"`
    Profile     UserProfile     `json:"profile"`
}

type Password struct {
    Plaintext *string
    Hash      string
}

type UserModel struct {
	DB *sql.DB
}

type UserID struct {
    Id uuid.UUID
}

var (
    ErrDuplicateEmail = errors.New("duplicate email")
)
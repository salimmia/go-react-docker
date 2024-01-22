package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/salimmia/go-architecture/internal/helpers"
	"github.com/salimmia/go-architecture/internal/models"
	"github.com/salimmia/go-architecture/internal/types"
	"github.com/salimmia/go-architecture/internal/utils"
)


func (m *Repository) RegistrationUser(w http.ResponseWriter, r *http.Request){
	var input struct {
        FirstName string `json:"first_name"`
        LastName  string `json:"last_name"`
		Email     string `json:"email"`
        Password  string `json:"password"`
    }
	
	err := helpers.ReadJSON(w, r, &input)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &models.User{
		Email: input.Email,
		FirstName: input.FirstName,
		LastName: input.LastName,
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	user.Password.Hash = hashedPassword
	user.Password.Plaintext = &input.Password
	
	userId, err := m.DB.RegistrationUser(user)
	if err != nil{
		switch {
        case errors.Is(err, models.ErrDuplicateEmail):
            http.Error(w, "A user with this email address already exists", http.StatusInternalServerError)
        default:
            http.Error(w, "Registration failed!", http.StatusInternalServerError)
        }
        return
	}

	user.ID = userId.Id
	user.CreatedAt = time.Now()
	
	err = helpers.WriteJSON(w, http.StatusOK, user, nil)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (m *Repository) UpdateUser(w http.ResponseWriter, r *http.Request){
	var userId models.UserID
	idParam := chi.URLParam(r, "user_id")

	id, err := uuid.Parse(idParam)
	if err != nil{
		http.Error(w, "could not parse user_id", http.StatusBadRequest)
		return
	}
	
	userId.Id = id
	db_user, err := m.DB.GetUserById(userId.Id)
	if err != nil {
		http.Error(w, "Errors", http.StatusInternalServerError)
		return
	}

	var input struct {
		FirstName   *string        `json:"first_name"`
		LastName    *string        `json:"last_name"`
		Thumbnail   *string        `json:"thumbnail"`
		PhoneNumber *string        `json:"phone_number"`
		BirthDate   types.NullTime `json:"birth_date"`
	}

	err = helpers.ReadJSON(w, r, &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if input.FirstName != nil {
		db_user.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		db_user.LastName = *input.LastName
	}
	if input.Thumbnail != nil {
		db_user.Thumbnail = input.Thumbnail
	}
	if input.PhoneNumber != nil {
		db_user.Profile.PhoneNumber = input.PhoneNumber
	}
	if input.BirthDate.Valid {
		db_user.Profile.BirthDate.Time = input.BirthDate.Time
	}

	update_user, err := m.DB.UpdateUser(db_user)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, update_user, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m *Repository) LogIn(w http.ResponseWriter, r *http.Request){
	var input struct{
		Email		string 		`json:"email"`
		Password 	string		`json:"password"`
	}

	err := helpers.ReadJSON(w, r, &input)

	if err != nil{
		return
	}

	user, err := m.DB.GetUserByEmail(input.Email)
	if err != nil{
		switch {
        case errors.Is(err, models.ErrRecordNotFound):
            http.Error(w, "Provide a valid email address", http.StatusBadRequest)
        default:
            http.Error(w, "login failed!", http.StatusInternalServerError)
        }
        return
	}
	
	loggedIn := utils.CheckPasswordHash(input.Password, user.Password.Hash)
	if !loggedIn{
		http.Error(w, "Password is not correct", http.StatusBadRequest)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, "Logged in susseccfully!", nil)
}

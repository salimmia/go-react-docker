package repository

type Repository interface{
	GetAllUsers() error
}
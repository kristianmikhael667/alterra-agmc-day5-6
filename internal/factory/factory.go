package factory

import (
	"main/database"
	"main/internal/repository"
)

type Factory struct {
	UsersRepository repository.Users
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		repository.NewUserRepository(db),
	}
}

package config

import (
	"tfg/user-service/pkg/domain/user"
	user2 "tfg/user-service/pkg/infraestructure/user"
	"tfg/user-service/pkg/libs/sql"
)

type Dependencies struct {
	UserRepository user.Repository
}

func BuildDependencies(config *Config) (*Dependencies, error) {
	db, err := sql.NewDB(config.DBConfig)

	if err != nil {
		return nil, err
	}
	userRepository := user2.NewPostgresUserRepository(db)
	return &Dependencies{
		userRepository,
	}, nil
}

package repository

import (
	"context"
	"log"
	"time"

	"authentication-service/entity"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	DBConn *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		DBConn: db,
	}
}

// create user details
func (repo *UserRepository) AddUserDetails(ctx context.Context, mConfigs *entity.User) error {
	res, err := repo.DBConn.Exec(ctx, `insert into users(email, first_name, last_name, password, created_at, updated_at) values ($1, $2, $3, $4, $5, $6)`, mConfigs.Email, mConfigs.FirstName, mConfigs.LastName, mConfigs.Password, time.Now(), time.Now())
	if err != nil {
		log.Fatalf("Unable to save user details. %v", err)
		return err
	}
	log.Println("Insert user", res)
	return nil
}

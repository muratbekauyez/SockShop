package repositories

import (
	"context"
	"Users/core"
	"Users/core/interfaces"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type UserRepository struct {
	pool pgxpool.Pool
}

func NewUserRepository(conn *pgxpool.Pool) interfaces.IUsersRepository {
	return &UserRepository{*conn}
}

func (r *UserRepository) CreateUser(user core.User) bool {
	sql := "INSERT INTO users(name, surname, username) VALUES($1, $2, $3) RETURNING id"
	row := r.pool.QueryRow(context.Background(), sql, user.Name, user.Surname, user.Username)
	var id int
	err := row.Scan(&id)
	if err != nil{
		log.Printf("Unable to Insert: %v\n", err)
		return false
	}
	return true

}

func (r *UserRepository) GetAllUsers() []*core.User {
	stmt := "SELECT * FROM users"
	rows, err := r.pool.Query(context.Background(), stmt)
	if err != nil {
		log.Fatal("Failed to SELECT:", err)
		return nil
	}
	defer rows.Close()
	users := []*core.User{}
	for rows.Next() {
		u := &core.User{}
		err = rows.Scan(&u.Id, &u.Name, &u.Surname, &u.Username)
		if err != nil {
			log.Fatalf("Failed to scan: %v", err)
			return nil
		}
		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	return users
}

func (r *UserRepository) GetUserById(id int) *core.User {
	stmt := "SELECT * FROM users WHERE id = $1"
	u := &core.User{}
	err := r.pool.QueryRow(context.Background(), stmt, id).Scan(&u.Id, &u.Name, &u.Surname, &u.Username)
	if err != nil {
		log.Println("Didn't find user with id ", id)
		return nil
	}
	return u
}

func (r *UserRepository) DeleteUser(user core.User) bool {
	_, err := r.pool.Exec(context.Background(), "DELETE FROM users WHERE id = $1", user.Id)
	if err != nil {
		return false
	}
	return true
}

func (r *UserRepository) UpdateUser(user core.User) bool {
	_, err := r.pool.Exec(context.Background(), "UPDATE users SET name = $1, surname = $2, username = $3 WHERE id = $4",
		user.Name, user.Surname, user.Username, user.Id)
	if err != nil {
		return false
	}
	return true
}
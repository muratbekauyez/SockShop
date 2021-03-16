package repositories

import (
	"Catalog/core"
	"Catalog/core/interfaces"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type SockRepository struct {
	pool pgxpool.Pool
}

func NewSockRepository(conn *pgxpool.Pool) interfaces.ISocksRepository {
	return &SockRepository{*conn}
}
func (r *SockRepository) CreateSock(sock core.Sock) bool {
	sql := "INSERT INTO socks(name, price, quantity) VALUES($1, $2, $3) RETURNING id"
	row := r.pool.QueryRow(context.Background(), sql, sock.Name, sock.Price, sock.Quantity)
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Printf("Unable to INSERT: %v\n", err)
		return false
	}
	return true
}

func (r SockRepository) GetAllSocks() []*core.Sock {
	stmt := "SELECT * FROM socks"
	rows, err := r.pool.Query(context.Background(), stmt)
	if err != nil {
		log.Fatal("Failed to SELECT:", err)
		return nil
	}
	defer rows.Close()
	socks := []*core.Sock{}
	for rows.Next() {
		s := &core.Sock{}
		err = rows.Scan(&s.Id, &s.Name, &s.Price, &s.Quantity)
		if err != nil {
			log.Fatalf("Failed to scan: %v", err)
			return nil
		}
		socks = append(socks, s)
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	return socks
}
func (r *SockRepository) GetSockById(id int) *core.Sock {
	stmt := "SELECT * FROM socks WHERE id = $1"
	s := &core.Sock{}
	err := r.pool.QueryRow(context.Background(), stmt, id).Scan(&s.Id, &s.Name, &s.Price, &s.Quantity)
	if err != nil {
		log.Println("Didn't find user with id ", id)
		return nil
	}
	return s
}

func (r *SockRepository) DeleteSock(sock core.Sock) bool {
	_, err := r.pool.Exec(context.Background(), "DELETE FROM socks WHERE id = $1", sock.Id)
	if err != nil {
		return false
	}
	return true
}
func (r SockRepository) UpdateSock(sock core.Sock) bool {
	_, err := r.pool.Exec(context.Background(), "UPDATE socks SET name = $1, price = $2, quantity = $3 WHERE id = $4",
		sock.Name, sock.Price, sock.Quantity, sock.Id)
	if err != nil {
		return false
	}
	return true
}

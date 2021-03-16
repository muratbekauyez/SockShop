package repositories

import (
	"Ordering/core"
	"Ordering/core/interfaces"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type OrderRepository struct {
	pool pgxpool.Pool
}

func NewOrderRepository(conn *pgxpool.Pool) interfaces.IOrdersRepository {
	return &OrderRepository{*conn}
}
func (r *OrderRepository) CreateOrder(order core.Order) bool {
	sql := "INSERT INTO orders(sum, destination) VALUES($1, $2) RETURNING id"
	row := r.pool.QueryRow(context.Background(), sql, order.Sum, order.Destination)
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Printf("Unable to INSERT: %v\n", err)
		return false
	}
	return true
}

func (r OrderRepository) GetAllOrders() []*core.Order {
	stmt := "SELECT * FROM orders"
	rows, err := r.pool.Query(context.Background(), stmt)
	if err != nil {
		log.Fatal("Failed to SELECT:", err)
		return nil
	}
	defer rows.Close()
	orders := []*core.Order{}
	for rows.Next() {
		o := &core.Order{}
		err = rows.Scan(&o.Id, &o.Sum, &o.Destination)
		if err != nil {
			log.Fatalf("Failed to scan: %v", err)
			return nil
		}
		orders = append(orders, o)
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	return orders
}
func (r *OrderRepository) GetOrderById(id int) *core.Order {
	stmt := "SELECT * FROM orders WHERE id = $1"
	o := &core.Order{}
	err := r.pool.QueryRow(context.Background(), stmt, id).Scan(&o.Id, &o.Sum, &o.Destination)
	if err != nil {
		log.Println("Didn't find order with id ", id)
		return nil
	}
	return o
}

func (r *OrderRepository) DeleteOrder(order core.Order) bool {
	_, err := r.pool.Exec(context.Background(), "DELETE FROM orders WHERE id = $1", order.Id)
	if err != nil {
		return false
	}
	return true
}
func (r OrderRepository) UpdateOrder(order core.Order) bool {
	_, err := r.pool.Exec(context.Background(), "UPDATE orders SET sum = $1, destination = $2 WHERE id = $3",
		order.Sum, order.Destination, order.Id)
	if err != nil {
		return false
	}
	return true
}

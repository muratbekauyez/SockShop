package main

import (
	"Ordering/pkg"
	"Ordering/pkg/repositories"
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func openDB(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Println("Connection for database is established")
		return nil, err
	}
	return pool, nil
}

func main() {
	dsn := flag.String("dsn", "postgresql://localhost/goFinal?user=postgres&password=qwerty", "PostGreSQL")
	flag.Parse()
	var err error
	pkg.Conn, err = openDB(*dsn)
	if err != nil{
		log.Fatalf("Failed to connect to db: ", err)
	}
	orderRepository = repositories.NewOrderRepository(pkg.Conn)
	router  := gin.Default()
	RouteOrders(router)
	router.Run(":4001") // listen and serve on port 4001
}


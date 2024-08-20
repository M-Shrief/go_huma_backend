package database

import (
	"context"
	"fmt"
	"go_huma_backend/internal/config"
	"os"

	"github.com/jackc/pgx/v5"
)

var Q *Queries

func Connect() (*pgx.Conn, error) {
	ctx := context.Background()

	connStr := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_NAME, config.DB_PASSWORD)
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		os.Exit(1)
	}
	fmt.Println("Database Connected")

	// err := db.Ping(context.Background())
	// if err != nil {
	// 	fmt.Printf("DB error: %v", err)
	// 	os.Exit(1)
	// }

	dataTypeNames := []string{
		"role",
		// An underscore prefix is an array type in pgtypes.
		"_role",
	}

	for _, typeName := range dataTypeNames {
		dataType, err := conn.LoadType(ctx, typeName)
		if err != nil {
			fmt.Printf("couldb't register database type: %v", err)
			os.Exit(1)
		}
		conn.TypeMap().RegisterType(dataType)
	}

	Q = New(conn)
	return conn, err
}

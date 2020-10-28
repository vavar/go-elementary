package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	conn, err := pgxpool.Connect(context.Background(), "postgresql://postgres:changeit@localhost")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	var sql string = `COPY temp FROM program '$1';`
	results, err := conn.Exec(context.Background(), sql, "ls -la")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	println(results)
}

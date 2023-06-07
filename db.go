package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DbInit() *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	var id int
	var title string
	var content string
	err = dbpool.QueryRow(context.Background(), "select * from posts;").Scan(&id, &title, &content)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(title, content)

	return dbpool
}

func NewPost(conn *pgxpool.Pool, title, content string) {
	_, err := conn.Exec(context.Background(), "INSERT INTO posts (title, content) VALUES ($1, $2)", title, content)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert into posts: %v\n", err)
	}

}

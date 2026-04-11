package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func initDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	createTable()
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		completed BOOLEAN DEFAULT FALSE
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func getTodos() ([]Todo, error) {
	rows, err := db.Query("SELECT id, title, completed FROM todos ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func createTodo(title string) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO todos (title) VALUES ($1) RETURNING id", title).Scan(&id)
	return id, err
}

func toggleTodo(id int, completed bool) error {
	_, err := db.Exec("UPDATE todos SET completed = $1 WHERE id = $2", completed, id)
	return err
}

func deleteTodo(id int) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}

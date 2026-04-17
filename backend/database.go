package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var db *sql.DB

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func initDB() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./data/todos.db"
	}

	// Adjust relative paths when running directly from the backend directory
	// (e.g. via `go run -C backend .`)
	if _, err := os.Stat("main.go"); err == nil {
		if _, err := os.Stat("../frontend"); err == nil {
			if dbPath == "./data/todos.db" {
				dbPath = "../data/todos.db"
			}
		}
	}

	// Ensure the directory for the database file exists
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("Failed to create database directory %s: %v", dir, err)
	}

	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Enable WAL mode for better concurrent read performance
	_, err = db.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		log.Printf("Warning: could not set WAL mode: %v", err)
	}

	fmt.Printf("📦 Database connected: %s\n", dbPath)
	createTable()
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
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
	result, err := db.Exec("INSERT INTO todos (title) VALUES (?)", title)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func toggleTodo(id int, completed bool) error {
	_, err := db.Exec("UPDATE todos SET completed = ? WHERE id = ?", completed, id)
	return err
}

func deleteTodo(id int) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}

package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mehmeddjug/goba/content/internal/core"
)

type ContentRepositorySqlite struct {
	db *sql.DB
}

func NewContentRepositorySqlite(db *sql.DB) *ContentRepositorySqlite {
	return &ContentRepositorySqlite{db: db}
}

/*
func (r *ContentRepositorySqlite) createTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        email TEXT
    );`
	_, err := r.db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	fmt.Println("Table created successfully")
}
*/

func (r *ContentRepositorySqlite) Create(title, username string) error {
	insertSQL := `INSERT INTO content (title, username) VALUES (?, ?)`
	result, err := r.db.Exec(insertSQL, title, username)
	if err != nil {
		log.Fatalf("Error inserting record: %v", err)
	}
	id, _ := result.LastInsertId()
	fmt.Printf("Record created with ID: %d\n", id)
	return err
}

func (r *ContentRepositorySqlite) Get(id int64) (*core.Content, error) {
	querySQL := `SELECT id, name, username FROM users WHERE id = ?`
	row := r.db.QueryRow(querySQL, id)
	var content core.Content
	err := row.Scan(&content.ID, &content.Title, &content.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No record found")
		} else {
			log.Fatalf("Error reading record: %v", err)
		}
	} else {
		fmt.Printf("Record: ID=%d, Title=%s, Username=%s\n", &content.ID, &content.Title, &content.Username)
	}
	return &content, nil
}

func (r *ContentRepositorySqlite) Update(id int64, name, email string) error {
	updateSQL := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	_, err := r.db.Exec(updateSQL, name, email, id)
	if err != nil {
		log.Fatalf("Error updating record: %v", err)
	}
	fmt.Println("Record updated successfully")
	return err
}

func (r *ContentRepositorySqlite) Delete(id int64) error {
	deleteSQL := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(deleteSQL, id)
	if err != nil {
		log.Fatalf("Error deleting record: %v", err)
	}
	fmt.Println("Record deleted successfully")
	return err
}

func (r *ContentRepositorySqlite) GetAll() ([]*core.Content, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var content []*core.Content
	for rows.Next() {
		var story core.Content
		if err := rows.Scan(&story.ID, &story.Title, &story.Username); err != nil {
			return nil, err
		}
		content = append(content, &story)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return content, nil
}

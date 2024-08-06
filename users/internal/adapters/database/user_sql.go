package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mehmeddjug/goba/users/internal/core"
)

type UserRepositorySqlite struct {
	db *sql.DB
}

func NewUserRepositorySqlite(db *sql.DB) *UserRepositorySqlite {
	return &UserRepositorySqlite{db: db}
}

/*
func (r *UserRepositorySqlite) createTable() {
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

func (r *UserRepositorySqlite) Create(name, email string) error {
	insertSQL := `INSERT INTO users (name, email) VALUES (?, ?)`
	result, err := r.db.Exec(insertSQL, name, email)
	if err != nil {
		log.Fatalf("Error inserting record: %v", err)
	}
	id, _ := result.LastInsertId()
	fmt.Printf("Record created with ID: %d\n", id)
	return err
}

func (r *UserRepositorySqlite) Get(id int64) (*core.User, error) {
	querySQL := `SELECT id, name, email FROM users WHERE id = ?`
	row := r.db.QueryRow(querySQL, id)
	var user core.User
	err := row.Scan(&user.ID, &user.Username, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No record found")
		} else {
			log.Fatalf("Error reading record: %v", err)
		}
	} else {
		fmt.Printf("Record: ID=%d, Name=%s, Email=%s\n", &user.ID, &user.Username, &user.Role)
	}
	return &user, nil
}

func (r *UserRepositorySqlite) Update(id int64, name, email string) error {
	updateSQL := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	_, err := r.db.Exec(updateSQL, name, email, id)
	if err != nil {
		log.Fatalf("Error updating record: %v", err)
	}
	fmt.Println("Record updated successfully")
	return err
}

func (r *UserRepositorySqlite) Delete(id int64) error {
	deleteSQL := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(deleteSQL, id)
	if err != nil {
		log.Fatalf("Error deleting record: %v", err)
	}
	fmt.Println("Record deleted successfully")
	return err
}

func (r *UserRepositorySqlite) GetAll() ([]*core.User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*core.User
	for rows.Next() {
		var user core.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

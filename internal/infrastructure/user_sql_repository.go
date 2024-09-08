package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mehmeddjug/goba/internal/domain"
)

type UserRepositorySql struct {
	db *sql.DB
}

func NewUserRepositorySql(db *sql.DB) *UserRepositorySql {
	return &UserRepositorySql{db: db}
}

func (r *UserRepositorySql) Create(username, email string) error {
	insertSQL := `INSERT INTO users (username, email) VALUES (?, ?)`
	result, err := r.db.Exec(insertSQL, username, email)
	if err != nil {
		log.Fatalf("Error inserting record: %v", err)
	}
	id, _ := result.LastInsertId()
	fmt.Printf("Record created with ID: %d\n", id)
	return err
}

func (r *UserRepositorySql) Read(id string) (*domain.User, error) {
	querySQL := `SELECT id, username, email FROM users WHERE id = ?`
	row := r.db.QueryRow(querySQL, id)
	var user domain.User
	err := row.Scan(&user.ID, &user.Username, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No record found")
		} else {
			log.Fatalf("Error reading record: %v", err)
		}
	} else {
		fmt.Printf("Record: ID=%s, Username=%s, Email=%s\n", user.ID, user.Username, user.Role)
	}
	return &user, nil
}

func (r *UserRepositorySql) Update(user *domain.User) error {
	updateSQL := `UPDATE users SET username = ?, email = ? WHERE id = ?`
	_, err := r.db.Exec(updateSQL, user.Username, user.Email, user.ID)
	if err != nil {
		log.Fatalf("Error updating record: %v", err)
	}
	fmt.Println("Record updated successfully")
	return err
}

func (r *UserRepositorySql) Delete(id string) error {
	deleteSQL := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(deleteSQL, id)
	if err != nil {
		log.Fatalf("Error deleting record: %v", err)
	}
	fmt.Println("Record deleted successfully")
	return err
}

package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mehmeddjug/goba/users/internal/core"
)

type ActivityLogRepositorySqlite struct {
	db *sql.DB
}

func NewActivityLogRepositorySqlite(db *sql.DB) *ActivityLogRepositorySqlite {
	return &ActivityLogRepositorySqlite{db: db}
}

func (r *ActivityLogRepositorySqlite) Create(aLog *core.ActivityLog) error {
	insertSQL := `INSERT INTO activity_logs (id, user_id, action, timestamp) VALUES (?, ?, ?, ?)`
	result, err := r.db.Exec(insertSQL, aLog.ID, aLog.UserID, aLog.Action, aLog.Timestamp)
	if err != nil {
		log.Fatalf("Error inserting record: %v", err)
	}
	id, _ := result.LastInsertId()
	fmt.Printf("Record created with ID: %d\n", id)
	return err
}

func (r *ActivityLogRepositorySqlite) List(userID string) ([]*core.ActivityLog, error) {
	rows, err := r.db.Query(`SELECT id, user_id, action, timestamp FROM activity_logs WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aLogs []*core.ActivityLog
	for rows.Next() {
		var aLog core.ActivityLog
		if err := rows.Scan(&aLog.ID, &aLog.UserID, &aLog.Action, &aLog.Timestamp); err != nil {
			return nil, err
		}
		aLogs = append(aLogs, &aLog)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return aLogs, nil
}

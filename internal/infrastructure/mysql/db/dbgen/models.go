// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package dbgen

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type TasksStatus string

const (
	TasksStatusNotStarted TasksStatus = "not_started"
	TasksStatusInProgress TasksStatus = "in_progress"
	TasksStatusDone       TasksStatus = "done"
	TasksStatusOnHold     TasksStatus = "on_hold"
	TasksStatusCanceled   TasksStatus = "canceled"
)

func (e *TasksStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TasksStatus(s)
	case string:
		*e = TasksStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for TasksStatus: %T", src)
	}
	return nil
}

type NullTasksStatus struct {
	TasksStatus TasksStatus `json:"tasks_status"`
	Valid       bool        `json:"valid"` // Valid is true if TasksStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTasksStatus) Scan(value interface{}) error {
	if value == nil {
		ns.TasksStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TasksStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTasksStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TasksStatus), nil
}

type Task struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      TasksStatus  `json:"status"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}

type User struct {
	ID             string    `json:"id"`
	Email          string    `json:"email"`
	Nickname       string    `json:"nickname"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	PasswordDigest string    `json:"password_digest"`
}

type UserTask struct {
	UserID string `json:"user_id"`
	TaskID string `json:"task_id"`
}

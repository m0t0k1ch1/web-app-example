// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package mysql

import (
	"database/sql/driver"
	"fmt"

	"app/gen/gqlgen"
	timeutil "github.com/m0t0k1ch1-go/timeutil/v4"
)

type TaskStatus string

const (
	TaskStatusUNCOMPLETED TaskStatus = "UNCOMPLETED"
	TaskStatusCOMPLETED   TaskStatus = "COMPLETED"
)

func (e *TaskStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TaskStatus(s)
	case string:
		*e = TaskStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for TaskStatus: %T", src)
	}
	return nil
}

type NullTaskStatus struct {
	TaskStatus TaskStatus
	Valid      bool // Valid is true if TaskStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTaskStatus) Scan(value interface{}) error {
	if value == nil {
		ns.TaskStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TaskStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTaskStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TaskStatus), nil
}

type Task struct {
	ID        uint64
	Title     string
	Status    gqlgen.TaskStatus
	UpdatedAt timeutil.Timestamp
	CreatedAt timeutil.Timestamp
}

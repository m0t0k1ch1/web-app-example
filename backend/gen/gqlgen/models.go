// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlgen

import (
	"fmt"
	"io"
	"strconv"

	"github.com/m0t0k1ch1-go/gqlutil"
)

type CompleteTaskError interface {
	IsCompleteTaskError()
}

type CreateTaskError interface {
	IsCreateTaskError()
}

type Error interface {
	IsError()
	GetMessage() string
}

type IConnection interface {
	IsIConnection()
	GetPageInfo() *PageInfo
	GetTotalCount() gqlutil.Int64
}

type Node interface {
	IsNode()
	GetId() string
}

type BadRequestError struct {
	Message string `json:"message"`
}

func (BadRequestError) IsError()                {}
func (this BadRequestError) GetMessage() string { return this.Message }

func (BadRequestError) IsCreateTaskError() {}

func (BadRequestError) IsCompleteTaskError() {}

type CompleteTaskInput struct {
	ClientMutationId *string `json:"clientMutationId,omitempty"`
	Id               string  `json:"id" validate:"required" en:"id"`
}

type CompleteTaskPayload struct {
	ClientMutationId *string           `json:"clientMutationId,omitempty"`
	Task             *Task             `json:"task,omitempty"`
	Error            CompleteTaskError `json:"error,omitempty"`
}

type CreateTaskInput struct {
	ClientMutationId *string `json:"clientMutationId,omitempty"`
	Title            string  `json:"title" validate:"min=1,max=32" en:"title"`
}

type CreateTaskPayload struct {
	ClientMutationId *string         `json:"clientMutationId,omitempty"`
	Task             *Task           `json:"task,omitempty"`
	Error            CreateTaskError `json:"error,omitempty"`
}

type Mutation struct {
}

type NoopInput struct {
	ClientMutationId *string `json:"clientMutationId,omitempty"`
}

type NoopPayload struct {
	ClientMutationId *string `json:"clientMutationId,omitempty"`
}

type PageInfo struct {
	EndCursor   *string `json:"endCursor,omitempty"`
	HasNextPage bool    `json:"hasNextPage"`
}

type Query struct {
}

type Task struct {
	Id     string     `json:"id"`
	Title  string     `json:"title"`
	Status TaskStatus `json:"status"`
}

func (Task) IsNode()            {}
func (this Task) GetId() string { return this.Id }

type TaskConnection struct {
	Edges      []*TaskEdge   `json:"edges"`
	PageInfo   *PageInfo     `json:"pageInfo"`
	TotalCount gqlutil.Int64 `json:"totalCount"`
}

func (TaskConnection) IsIConnection()                    {}
func (this TaskConnection) GetPageInfo() *PageInfo       { return this.PageInfo }
func (this TaskConnection) GetTotalCount() gqlutil.Int64 { return this.TotalCount }

type TaskEdge struct {
	Cursor string `json:"cursor"`
	Node   *Task  `json:"node"`
}

type TaskStatus string

const (
	TaskStatusUncompleted TaskStatus = "UNCOMPLETED"
	TaskStatusCompleted   TaskStatus = "COMPLETED"
)

var AllTaskStatus = []TaskStatus{
	TaskStatusUncompleted,
	TaskStatusCompleted,
}

func (e TaskStatus) IsValid() bool {
	switch e {
	case TaskStatusUncompleted, TaskStatusCompleted:
		return true
	}
	return false
}

func (e TaskStatus) String() string {
	return string(e)
}

func (e *TaskStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskStatus", str)
	}
	return nil
}

func (e TaskStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

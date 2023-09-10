package model

import "time"

type Todo struct {
	ID			int			`json:"id"`
	Name		string		`json:"name"`
	Completed	bool		`json:"completed"`
	CreatedAt	time.Time	`json:"created_at"`
}

type DBHandler interface {
	GetTodos() []*Todo
	AddTodo(name string) *Todo
	RemoveTodo(id int) bool
	UpdateTodo(id int, completed bool) bool
	Close()
}

// NOTE: 패키지가 시작될 때 한 번만 호출되는 함수
var handler DBHandler
func NewDBHandler(filepath string) DBHandler {
	return newSqliteHandler(filepath)
}
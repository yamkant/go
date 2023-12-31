package model

import "time"

type Todo struct {
	ID			int			`json:"id"`
	Name		string		`json:"name"`
	Completed	bool		`json:"completed"`
	CreatedAt	time.Time	`json:"created_at"`
}

type DBHandler interface {
	GetTodos(sessionId string) []*Todo
	AddTodo(name string, sessionId string) *Todo
	RemoveTodo(id int) bool
	UpdateTodo(id int, completed bool) bool
	Close()
}

// NOTE: 패키지가 시작될 때 한 번만 호출되는 함수
var handler DBHandler
func NewDBHandler(dbConn string, reset bool) DBHandler {
	// return newSqliteHandler(dbConn, reset)
	return newPQHandler(dbConn, reset)
}
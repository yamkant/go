package model

import "time"

type Todo struct {
	ID			int		`json:"id"`
	Name		string		`json:"name"`
	Completed	bool		`json:"completed"`
	CreatedAt	time.Time	`json:"created_at"`
}

type dbHandler interface {
	getTodos() []*Todo
	addTodo(name string) *Todo
	removeTodo(id int) bool
	updateTodo(id int, completed bool) bool
}

// NOTE: 패키지가 시작될 때 한 번만 호출되는 함수
var handler dbHandler
func init() {
	// NOTE: handler의 구현체를 생성
	handler = newMemoryHandler()
}

func GetTodos() []*Todo {
	return handler.getTodos()
}

func AddTodo(name string) *Todo {
	return handler.addTodo(name)
} 

func RemoveTodo(id int) bool {
	return handler.removeTodo(id)
}

func UpdateTodo(id int, completed bool) bool {
	return handler.updateTodo(id, completed)
}
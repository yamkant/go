package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"example.com/m/model"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := model.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	todo := model.AddTodo(name)
	rd.JSON(w, http.StatusCreated, todo)
}

type Success struct {
	Success bool `json:"success"`
}

type TodoCompleted struct {
	Completed	bool		`json:"completed"`
}

func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := model.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var data TodoCompleted
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}

	ok := model.UpdateTodo(id, data.Completed)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func MakeNewHandler() http.Handler {

	rd = render.New()
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	router.HandleFunc("/todos", addTodoHandler).Methods("POST")
	router.HandleFunc("/todos/{id:[0-9]+}", updateTodoHandler).Methods("PATCH")
	router.HandleFunc("/todos/{id:[0-9]+}", removeTodoHandler).Methods("DELETE")

	return router
}
package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"example.com/m/model"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("UUID_SESSION_KEY")))
var rd *render.Render = render.New()


type AppHandler struct {
	http.Handler
	db model.DBHandler
}

var getSessionID = func (r *http.Request) string {
	session, err := store.Get(r, "session")
	if err != nil {
		return ""
	}

	// Set some session values.
	val := session.Values["id"]
	if val == nil {
		return ""
	}
	return val.(string)
}

func (a *AppHandler)indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func  (a *AppHandler)getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getSessionID(r)
	list := a.db.GetTodos(sessionId)
	rd.JSON(w, http.StatusOK, list)
}

func  (a *AppHandler)addTodoHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getSessionID(r)
	name := r.FormValue("name")
	todo := a.db.AddTodo(name, sessionId)
	rd.JSON(w, http.StatusCreated, todo)
}

type Success struct {
	Success bool `json:"success"`
}

type TodoCompleted struct {
	Completed bool `json:"completed"`
}

func  (a *AppHandler)updateTodoHandler(w http.ResponseWriter, r *http.Request) {
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

	ok := a.db.UpdateTodo(id, data.Completed)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func  (a *AppHandler)removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := a.db.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}
func (a *AppHandler) Close() {
	a.db.Close()
}

func CheckSignin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	loginURL := "/signin"
	oauthURL := "/auth"
	if strings.Contains(r.URL.Path, loginURL) || strings.Contains(r.URL.Path, oauthURL){
		next(w, r)
	}
	sessionID := getSessionID(r)
	if sessionID != "" {
		next(w, r)
	}

	http.Redirect(w, r, "/signin.html", http.StatusTemporaryRedirect)
}

func MakeNewHandler(filepath string, reset bool) *AppHandler {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewRecovery(), 
		negroni.NewLogger(),
		negroni.HandlerFunc(CheckSignin),
		negroni.NewStatic(http.Dir("public")),
	)

	// NOTE: middleware 역할
	n.UseHandler(r)

	a := &AppHandler{
		Handler: n,
		db: model.NewDBHandler(filepath, reset),
	}

	r.HandleFunc("/", a.indexHandler)
	r.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", a.updateTodoHandler).Methods("PATCH")
	r.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")

	r.HandleFunc("/auth/google/login", googleLoginHandler)
	r.HandleFunc("/auth/google/callback", googleAuthCallback)

	return a
}
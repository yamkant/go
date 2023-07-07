package myapp

import (
	"fmt"
    "encoding/json"
	"strconv"
	"time"
	"net/http"
    "github.com/gorilla/mux"
)


type User struct {
    ID                  int
    FirstName           string `json:"first_name"`
    LastName            string `json:"last_name"`
    Email               string `json:"email"`
    CreatedAt           time.Time `json:"created_at"`
}

var userMap map[int]*User
var lastID int

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello World")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Get UserInfo by /users/{id}")
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprint(w, err)
        return
    }

    user, ok := userMap[id]
    if !ok {
        w.WriteHeader(http.StatusOK)
        fmt.Fprint(w, "No User ID:", id)
        return
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    data, _ := json.Marshal(user)
    fmt.Fprint(w, string(data))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
    user := new(User)
    err := json.NewDecoder(r.Body).Decode(user)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprint(w, err)
        return
    }
    lastID++
    user.ID = lastID
    user.CreatedAt = time.Now()
    userMap[user.ID] = user


    w.WriteHeader(http.StatusCreated)
    data, _ := json.Marshal(user)
    fmt.Fprint(w, string(data))
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprint(w, err)
        return
    }

    // NOTE: Map에 없는 경우
    _, ok := userMap[id]
    if !ok {
        w.WriteHeader(http.StatusOK)
        fmt.Fprint(w, "No User ID:", id)
        return
    }
    delete(userMap, id)
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Deleted User ID:", id)
}

func NewHandler() http.Handler {
    userMap = make(map[int]*User)
    mux := mux.NewRouter()

    mux.HandleFunc("/", indexHandler)
    mux.HandleFunc("/users", usersHandler).Methods("GET")
    mux.HandleFunc("/users", createUserHandler).Methods("POST")
    mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler).Methods("GET")
    mux.HandleFunc("/users/{id:[0-9]+}", deleteUserHandler).Methods("DELETE")

    return mux
}
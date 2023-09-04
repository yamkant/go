package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

type User struct {
	Name	string `json:"name"`
	Email	string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "tester", Email: "tester@example.com"}

	rd.JSON(w, http.StatusOK, user)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		rd.Text(w, http.StatusBadRequest, err.Error())
		return 
	}
	user.CreatedAt = time.Now()

	rd.JSON(w, http.StatusCreated, user)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "tester", Email: "tester@example.com"}

	rd.HTML(w, http.StatusOK, "body", user)
}

func main() {
	rd = render.New(render.Options{
		Directory: "templates",
		Extensions: []string{ ".html", ".tmpl", },
		Layout: "hello",
	}) // html 확장자도 읽을 수 있도록 설정. rd.HTML 기본 확장자는 .tmpl
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	// mux.Handle("/", http.FileServer(http.Dir("public")))
	// NOTE: negroni - logger 기능 기본 제공 및 static file 서빙 (public 폴더 기본)
	n := negroni.Classic()
	n.UseHandler(mux)
	http.ListenAndServe(":3000", n)
}
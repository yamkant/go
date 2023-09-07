package main

import (
	"log"
	"net/http"

	"example.com/m/app"

	"github.com/urfave/negroni"
)

func main() {
	mux := app.MakeNewHandler()
	n := negroni.Classic()
	n.UseHandler(mux)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
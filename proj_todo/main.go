package main

import (
	"log"
	"net/http"

	"example.com/m/app"
)

func main() {
	mux := app.MakeNewHandler("./test.db")
	defer mux.Close()


	log.Println("Started App")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}
}
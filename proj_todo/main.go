package main

import (
	"log"
	"net/http"
	"os"

	"example.com/m/app"
)

func main() {
	mux := app.MakeNewHandler(os.Getenv("DATABASE_URL"))
	defer mux.Close()


	log.Println("Started App")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}
}
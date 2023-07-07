package main

import (
	"net/http"
	"example.com/user/web/myapp"
)

func main() {
    http.ListenAndServe(":3000", myapp.NewHandler())
}
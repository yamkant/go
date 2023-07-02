package main

import (
	"net/http"

	"github.com/yamkant/go/web/myapp"
)

func main() {
    http.ListenAndServe(":3000", myapp.NewHttpHandler)
}
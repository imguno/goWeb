package main

import (
	"net/http"

	"github.com/imguno/goWeb/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHTTPHandler())
}

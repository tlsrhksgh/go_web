package main

import (
	"myapp/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(
		":3000",
		myapp.NewHttpHandler(),
	)
}

package main

import (
	"net/http"

	"github.com/dancarlocanlas/go-getting-started/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

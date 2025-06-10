package main

import (
	"net/http"

	"github.com/riedel-gabriela/bootcamp-meli/goweb/aula10get/cmd/http/router"
)

func main() {
	r := router.NewRouter()
	if err := http.ListenAndServe(":3000", r.MapRoutes()); err != nil {
		panic(err)
	}
}

//TODO implementar getid, get by param, e post.

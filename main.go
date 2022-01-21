package main

import (
	"net/http"

	"github.com/AndreCDiniz/aplicacaoweb/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}

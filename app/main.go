package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/router/rotas"
)

func main() {
	fmt.Print("Rodando App!")

	r := router.Gerar()
	rotas.Configurar(r)
	log.Fatalln(http.ListenAndServe(":3000", r))

}

package rotas

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Rota struct {
	URI string
	Metodo string
	Funcao func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}

func Configurar(r *mux.Router) *mux.Router   {
	rotas := rotasUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}

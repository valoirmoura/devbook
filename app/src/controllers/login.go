package controllers

import "net/http"

// CarreagarTelaDeLogin vai renderizar a Tela de Login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tela de Login!"))
}

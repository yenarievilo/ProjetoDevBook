package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// escreve infos da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// Verifica se tem dados dentro dos cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, erro := cookies.Ler(r); erro != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}

		proximaFuncao(w, r)
	}
}

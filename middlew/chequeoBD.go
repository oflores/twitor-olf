package middlew

import (
	"net/http"

	"github.com/oflores/twitor-olf/bd"
)

/* Chequeo BD middleware que permite conocer esstao de DB */
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la Base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}

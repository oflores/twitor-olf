package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/oflores/twitor-olf/middlew"
	"github.com/oflores/twitor-olf/routers"
	"github.com/rs/cors"
)

/* se usa para manejar el puero handler y listeners*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/Registro", middlew.ChequeoDB(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

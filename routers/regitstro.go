package routers

import (
	"encoding/json"
	"net/http"

	"github.com/oflores/twitor-olf/bd"
	"github.com/oflores/twitor-olf/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "El password debe ser al menos de 6 caracteres", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya hay un usuario registrado con ese email", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar el registro"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No fue posible insertar el nuevo registro", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

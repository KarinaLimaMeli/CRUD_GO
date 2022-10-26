package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func List(w http.Response, r * http.Request){
todos, err := models.GetAll()
if err != nil {
	log.Printf("Erro ao obter registros: %v", err)
}
	w.Header().Add("Content-Type", "application/jason")
	json.NewEncoder(w).Encode(todos)

}
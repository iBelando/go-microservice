package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iBelando/go-microservice/commons"
	"github.com/iBelando/go-microservice/models"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	personas := []models.Persona{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&personas)

	json, _ := json.Marshal(personas)

	commons.SendResponse(w, http.StatusOK, json)
}

func Get(w http.ResponseWriter, r *http.Request) {
	persona := models.Persona{}

	id := mux.Vars(r)["id"]

	db := commons.GetConnection()
	defer db.Close()

	db.Find(&persona, id)

	if persona.ID > 0 {
		json, _ := json.Marshal(persona)

		commons.SendResponse(w, http.StatusOK, json)
	} else {
		commons.SendError(w, http.StatusNotFound)
	}
}

func Save(w http.ResponseWriter, r *http.Request) {
	persona := models.Persona{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(r.Body).Decode(persona)

	if error != nil {
		log.Fatal(error)
		commons.SendError(w, http.StatusBadRequest)
		return
	}

	error = db.Create(&persona).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(w, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(persona)

	commons.SendResponse(w, http.StatusCreated, json)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	persona := models.Persona{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(r)["id"]

	db.Find(&persona, id)

	if persona.ID > 0 {
		db.Delete(persona)
		commons.SendResponse(w, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(w, http.StatusNotFound)
	}
}

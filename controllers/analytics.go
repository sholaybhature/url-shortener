package controllers

import (
	"encoding/json"
	"net/http"
	"url-shorten/models"

	"github.com/julienschmidt/httprouter"
)

func GetURLAnalytics(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id := r.FormValue("id")
	if id != "" {
		obj, err := models.GetURLFromDB(id)
		if err != nil {
			http.Error(w, "try again later", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(obj)
	}
}

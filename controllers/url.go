package controllers

import (
	"encoding/json"
	"net/http"
	"url-shorten/utils"

	"github.com/julienschmidt/httprouter"
)

type sendResponse struct {
	Link      string
	ShortLink string
}

type requestBodyLink struct {
	URL *string `json:"link"`
}

func CreateShortenedURL(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	bodyLink := requestBodyLink{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&bodyLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if bodyLink.URL == nil {
		http.Error(w, "'link' field is missing", http.StatusBadRequest)
		return
	}
	if flag := utils.IsUrl(*bodyLink.URL); !flag {
		http.Error(w, "'link' is not a valid url", http.StatusBadRequest)
		return
	}
	encodedURL := utils.EncodeURL(*bodyLink.URL)
	res := sendResponse{
		Link:      *bodyLink.URL,
		ShortLink: encodedURL,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

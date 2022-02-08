package controllers

import (
	"encoding/json"
	"net/http"
	"url-shorten/models"
	"url-shorten/utils"

	"github.com/julienschmidt/httprouter"
)

type sendResponse struct {
	Link      string `json:"link"`
	ShortLink string `json:"shortLink"`
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
	// run for loop till maxlimit if there's a collision?
	err = models.SaveURLToDB(encodedURL, *bodyLink.URL)
	if err != nil {
		http.Error(w, "try again", http.StatusBadRequest)
		return
	}

	res := sendResponse{
		Link:      *bodyLink.URL,
		ShortLink: r.Host + "/" + encodedURL,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

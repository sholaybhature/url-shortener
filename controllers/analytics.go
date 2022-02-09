package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"url-shorten/models"

	"github.com/julienschmidt/httprouter"
)

func GetURLAnalytics(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id := r.FormValue("id")
	device := r.FormValue("device")
	fmt.Println(id, device)
	if id != "" {
		obj, err := models.GetURLFromDB(id)
		if err != nil {
			http.Error(w, "try again later", http.StatusInternalServerError)
			return
		}
		if device != "" {
			// deep copy manually
			copyObj := models.URLObj{Id: obj.Id, URL: obj.URL, Count: obj.Count}
			// copyObj.Visitors = nil
			fmt.Println(copyObj, obj)
			for i, val := range obj.Visitors {
				fmt.Println(val.Ip)
				fmt.Println(strings.ToLower(val.Device))
				fmt.Println(strings.ToLower(device))
				if strings.Contains(strings.ToLower(val.Device), strings.ToLower(device)) {
					fmt.Println(1)
					copyObj.Visitors = append(copyObj.Visitors, val)
					fmt.Print(i, val)
				}
			}
			json.NewEncoder(w).Encode(copyObj)
		} else {
			json.NewEncoder(w).Encode(obj)
		}
	} else {
		http.Error(w, "no id provided", http.StatusNotFound)
		return
	}

}

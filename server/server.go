package server

import (
	"encoding/json"
	"net/http"

	"github.com/akbarismail/rest/model"
)

func Students(w http.ResponseWriter, r *http.Request) {
	data := []model.Student{
		{Name: "glog", Age: 20},
		{Name: "bruno", Age: 19},
		{Name: "alu", Age: 25},
	}

	jsonInBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonInBytes)
}

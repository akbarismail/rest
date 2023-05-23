package client

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"path"

	"github.com/akbarismail/rest/model"
)

func Student(w http.ResponseWriter, r *http.Request) {
	filepath := path.Join("template", "student.html")

	parseFile, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// request data student dari client ke server
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://localhost:8000/api/student", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// membaca respon body dari server
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var students []model.Student
	err = json.Unmarshal(body, &students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title": "Belajar REST API Golang",
		"data":  students,
	}

	// render data ke template
	err = parseFile.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

package upload

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
)

type FormData struct {
	file      multipart.File
	cloudName string
	apiKey    string
	apiSecret string
}

type errResponse struct {
	Error        string   `json:"error"`
	ErrorMessage error    `json:"errorMessage,omitempty"`
	Fields       []string `json:"fields,omitempty"`
}

func Upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	file, header, _ := r.FormFile("img")

	if file == nil {
		err := errResponse{Error: "No image was found in your request."}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(err)
		return
	}

	data := FormData{
		file:      file,
		cloudName: r.FormValue("cloudName"),
		apiKey:    r.FormValue("cloudApiKey"),
		apiSecret: r.FormValue("cloudApiSecret"),
	}
	handleEmptyFields(data, w)

	res, err := uploadImage(file, header, data)

	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(errResponse{Error: "An error ocurred while trying to upload your image.", ErrorMessage: err})
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
}

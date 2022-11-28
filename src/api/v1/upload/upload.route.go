package upload

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
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

	// Iterate all fields from data struct and append every empty field inside emptyFields var
	values := reflect.ValueOf(data)
	var emptyFields []string
	types := values.Type()

	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).String() == "" {
			emptyFields = append(emptyFields, types.Field(i).Name)
		}
		fmt.Println(values.Field(i).String())
	}

	if len(emptyFields) > 0 {
		err := errResponse{Error: "There is one or more fields that are empty:", Fields: emptyFields}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(err)
		return
	}

	ctx := context.Background()
	filenameWithoutExtension := strings.Split(header.Filename, ".")[0]
	cld, _ := cloudinary.NewFromParams(data.cloudName, data.apiKey, data.apiSecret)
	res, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: filenameWithoutExtension})

	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(errResponse{Error: "An error ocurred while trying to upload your image.", ErrorMessage: err})
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
}

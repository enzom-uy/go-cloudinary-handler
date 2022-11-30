package upload

import (
	"context"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func handleEmptyFields(data FormData, w http.ResponseWriter) []string {
	// Iterate all fields from data struct and append every empty field inside emptyFields var
	values := reflect.ValueOf(data)
	var emptyFields []string
	types := values.Type()

	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).String() == "" {
			emptyFields = append(emptyFields, types.Field(i).Name)
		}
	}
	return emptyFields
}

func uploadImage(file multipart.File, header *multipart.FileHeader, data FormData) (*uploader.UploadResult, error) {
	ctx := context.Background()
	filenameWithoutExtension := strings.Split(header.Filename, ".")[0]
	cld, _ := cloudinary.NewFromParams(data.cloudName, data.apiKey, data.apiSecret)
	res, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: filenameWithoutExtension})

	return res, err
}

package upload

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	file, header, _ := r.FormFile("img")
	cldCloudName := r.FormValue("cloudName")
	cldApiKey := r.FormValue("cloudApiKey")
	cldApiSecret := r.FormValue("cloudApiSecret")
	w.Header().Set("Content-Type", "application/json")

	if cldCloudName == "" || cldApiKey == "" || cldApiSecret == "" || file == nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("There are some fields missing in the request.")
		return
	}

	ctx := context.Background()
	filenameWithoutExtension := strings.Split(header.Filename, ".")[0]
	cld, _ := cloudinary.NewFromParams(cldCloudName, cldApiKey, cldApiSecret)
	res, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: filenameWithoutExtension})
	fmt.Println(err)

	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("There was an error while trying to upload the image: ")
		json.NewEncoder(w).Encode(err)
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)

}

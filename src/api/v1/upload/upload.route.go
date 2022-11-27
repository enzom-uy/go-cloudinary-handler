package upload

import (
	"net/http"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("img")
}

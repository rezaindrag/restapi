package helper

import (
	"encoding/json"
	"net/http"
)

// JSON return json encode
func JSON(w http.ResponseWriter, msg interface{}, httpStatus int) {
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(msg)
}

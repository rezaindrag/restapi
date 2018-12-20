package helper

import (
	"encoding/json"
	"net/http"

	"github.com/rezaindrag/restapi/api/structs"
)

// JSON return json encode
func JSON(w http.ResponseWriter, errorMessege structs.ErrorMsg) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(errorMessege)
}

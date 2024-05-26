package helpers

import (
	"encoding/json"
	"net/http"
)

func DecodeBody(w http.ResponseWriter, r *http.Request, payload interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		Response(w, http.StatusBadRequest, err.Error(), nil)
		return err
	}
	return nil
}

package helpers

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic(err)
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		panic(err)
	}
}

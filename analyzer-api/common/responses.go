package common

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data     string
	Status   bool
	Messsage interface{}
}

// RespondWithError return error message
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondwithJSON(w, code, map[string]string{"error": msg})
}

// RespondwithJSON write json response format
func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	t := payload
	s := true
	var m interface{}

	if code != 200 {
		s = false
		t = nil
		m = payload
	}

	response, _ := json.Marshal(map[string]interface{}{
		"data":    t,
		"status":  s,
		"message": m,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

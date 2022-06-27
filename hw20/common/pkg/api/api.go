package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
	Code  int         `json:"-"`
}

func RenderJSON(w http.ResponseWriter, resp *Response) {
	log.Printf("Response: %+v", resp)

	payload, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.Code)
	w.Write(payload)
}

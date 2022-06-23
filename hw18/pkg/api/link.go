package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"net/http"
)

func (a *Api) addLink(w http.ResponseWriter, r *http.Request) {
	resp := &Response{Code: http.StatusOK}

	id, err := gonanoid.New(8)
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = "server error"
		renderJSON(w, resp)
		return
	}

	params := make(map[string]string)
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		resp.Code = http.StatusUnprocessableEntity
		resp.Error = "bad params"
		renderJSON(w, resp)
		return
	}

	a.Lock()
	a.data[id] = params["url"]
	a.Unlock()

	resp.Data = fmt.Sprintf("https://%s/%s", a.addr, id)
	renderJSON(w, resp)
}

func (a *Api) showLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resp := &Response{Code: http.StatusOK}

	a.Lock()
	if url, ok := a.data[vars["id"]]; ok {
		resp.Data = url
		renderJSON(w, resp)
		return
	}
	a.Unlock()

	resp.Code = http.StatusNotFound
	resp.Error = "URL not found"
	renderJSON(w, resp)
}

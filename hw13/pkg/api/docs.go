package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kevin-glare/hardcode-dev-go/hw5/pkg/crawler"
	"net/http"
	"strconv"
)

func (api *Api) index(w http.ResponseWriter, r *http.Request) {
	resp := &Response{Code: http.StatusOK}
	query := r.URL.Query().Get("q")

	if len(query) > 0 {
		resp.Data = api.store.FindByQuery(query)
	} else {
		resp.Data = api.store.Docs
	}

	renderJSON(w, resp)
}

func (api *Api) create(w http.ResponseWriter, r *http.Request) {
	resp := &Response{Code: http.StatusOK}
	var doc crawler.Document

	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		resp.Code = http.StatusBadRequest
		renderJSON(w, resp)
		return
	}

	api.store.Lock()
	doc.ID = api.store.Docs[len(api.store.Docs)-1].ID + 1
	api.store.Docs = append(api.store.Docs, doc)
	api.store.Unlock()

	resp.Data = doc
	resp.Code = http.StatusCreated
	renderJSON(w, resp)
}

func (api *Api) read(w http.ResponseWriter, r *http.Request) {
	resp := &Response{Code: http.StatusOK}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		resp.Error = err.Error()
		resp.Code = http.StatusBadRequest
		renderJSON(w, resp)
		return
	}

	doc := api.store.FindByID(id)
	if doc == nil {
		resp.Error = fmt.Sprintf("id %d not found\n", id)
		resp.Code = http.StatusNotFound
		renderJSON(w, resp)
		return
	}

	resp.Data = &doc
	renderJSON(w, resp)
}

func (api *Api) update(w http.ResponseWriter, r *http.Request) {
	resp := &Response{Code: http.StatusOK}
	var doc crawler.Document

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		resp.Error = err.Error()
		resp.Code = http.StatusBadRequest
		renderJSON(w, resp)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		resp.Code = http.StatusBadRequest
		renderJSON(w, resp)
		return
	}

	api.store.Lock()
	doc.ID = id
	for i, d := range api.store.Docs {
		if d.ID == doc.ID {
			api.store.Docs[i] = doc
			break
		}
	}
	api.store.Unlock()

	resp.Data = doc
	renderJSON(w, resp)
}

func (api *Api) destroy(w http.ResponseWriter, r *http.Request) {
	resp := &Response{Code: http.StatusOK}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		resp.Error = err.Error()
		resp.Code = http.StatusBadRequest
		renderJSON(w, resp)
		return
	}

	api.store.Lock()
	for i, doc := range api.store.Docs {
		if doc.ID == id {
			api.store.Docs = append(api.store.Docs[:i], api.store.Docs[i+1:]...)
			break
		}
	}
	api.store.Unlock()

	renderJSON(w, resp)
}

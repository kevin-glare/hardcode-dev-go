package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Api struct {
	router *mux.Router
	store  *Store
}

type Response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
	Code  int         `json:"-"`
}

func Run(addr string) {
	api := &Api{
		router: mux.NewRouter(),
		store:  NewStore(),
	}
	api.endpoints()
	api.router.Use(contentTypeMiddleware)

	http.Handle("/", api.router)
	srv := &http.Server{
		Handler:      api.router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func (api *Api) endpoints() {
	api.router.HandleFunc("/api/v1/docs", api.index).Methods(http.MethodGet)
	api.router.HandleFunc("/api/v1/docs", api.create).Methods(http.MethodPost)
	api.router.HandleFunc("/api/v1/docs/{id}", api.read).Methods(http.MethodGet)
	api.router.HandleFunc("/api/v1/docs/{id}", api.update).Methods(http.MethodPut)
	api.router.HandleFunc("/api/v1/docs/{id}", api.destroy).Methods(http.MethodDelete)
}

func renderJSON(w http.ResponseWriter, resp *Response) {
	log.Printf("Response: %+v", resp)

	payload, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.Code)
	w.Write(payload)
}

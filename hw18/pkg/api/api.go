package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
	"time"
)

type Api struct {
	router *mux.Router
	addr   string
	sync.Mutex
	data map[string]string
}

type Response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
	Code  int         `json:"-"`
}

func Run(addr string) {
	api := &Api{
		router: mux.NewRouter(),
		data:   make(map[string]string),
		addr:   addr,
	}

	api.endpoints()

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
	api.router.HandleFunc("/api/v1/links", api.addLink).Methods(http.MethodPost)
	api.router.HandleFunc("/api/v1/links/{id}", api.showLink).Methods(http.MethodGet)
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

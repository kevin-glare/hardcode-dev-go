package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kevin-glare/hardcode-dev-go/hw19/analytic_service/pkg/kfk"
	"log"
	"net/http"
	"sync"
	"time"
)

type Api struct {
	router *mux.Router
	addr   string
	kfk    *kfk.Analytic

	sync.Mutex
	data map[string]string
}

type Response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
	Code  int         `json:"-"`
}

func Run(addr string, kfk *kfk.Analytic) {
	api := &Api{
		router: mux.NewRouter(),
		addr:   addr,
		kfk:    kfk,
		data:   make(map[string]string),
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

func (a *Api) endpoints() {
	a.router.HandleFunc("/tt/{id}", a.redirectToLink).Methods(http.MethodGet)
	a.router.HandleFunc("/api/v1/links", a.addLink).Methods(http.MethodPost)
	a.router.HandleFunc("/api/v1/links/{id}", a.showLink).Methods(http.MethodGet)
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

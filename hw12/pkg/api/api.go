package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	r "github.com/kevin-glare/hardcode-dev-go/hw12/pkg/repository"
	"github.com/kevin-glare/hardcode-dev-go/hw5/pkg/crawler"
	"log"
	"net/http"
	"time"
)

var repository = r.NewRepository()

func Run(port int) {
	router := mux.NewRouter()
	endpoints(router)
	http.Handle("/", router)

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%v", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func endpoints(router *mux.Router) {
	router.HandleFunc("/index", indexHandler)
	router.HandleFunc("/docs", docsHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "index")
}

func docsHandler(w http.ResponseWriter, r *http.Request) {
	var response []crawler.Document
	code := http.StatusOK

	query := r.URL.Query().Get("q")

	if len(query) > 0 {
		response = repository.Find(query)
	} else {
		response = repository.Docs
	}

	payload, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		code = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(payload)
}

package api

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/delivery"
	"github.com/kevin-glare/hardcode-dev-go/hw20/link-service/pkg/link"
	"log"
	"net/http"
	"time"
)

type Api struct {
	router *mux.Router
	addr   string
	l      *link.Struct
}

func Run(addr string, link *link.Struct) {
	api := &Api{
		router: mux.NewRouter(),
		addr:   addr,
		l:      link,
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
	a.router.HandleFunc("/api/v1/links", a.newLink).Methods(http.MethodPost)
	a.router.HandleFunc("/api/v1/links/{shortURL}", a.link).Methods(http.MethodGet)
}

func (a *Api) link(w http.ResponseWriter, r *http.Request) {
	resp := &delivery.Response{Code: http.StatusOK}
	vars := mux.Vars(r)

	link, err := a.l.Link(context.Background(), vars["shortURL"])
	if err != nil {
		resp.Code = http.StatusNotFound
		resp.Error = err.Error()
		delivery.RenderJSON(w, resp)
		return
	}

	resp.Data = link.Url
	delivery.RenderJSON(w, resp)
}

func (a *Api) newLink(w http.ResponseWriter, r *http.Request) {
	resp := &delivery.Response{Code: http.StatusOK}
	params := make(map[string]string)

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		resp.Code = http.StatusUnprocessableEntity
		resp.Error = err.Error()
		delivery.RenderJSON(w, resp)
		return
	}

	shortLink, err := a.l.NewLink(context.Background(), params["url"])
	if err != nil {
		resp.Code = http.StatusUnprocessableEntity
		resp.Error = err.Error()
		delivery.RenderJSON(w, resp)
		return
	}

	resp.Data = shortLink
	delivery.RenderJSON(w, resp)
}

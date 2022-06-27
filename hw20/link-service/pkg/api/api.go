package api

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/api"
	"github.com/kevin-glare/hardcode-dev-go/hw20/link-service/pkg/service"
	"log"
	"net/http"
	"time"
)

type Api struct {
	router  *mux.Router
	addr    string
	service *service.LinkService
}

func Run(addr string, linkService *service.LinkService) {
	api := &Api{
		router:  mux.NewRouter(),
		service: linkService,
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
	resp := &api.Response{Code: http.StatusOK}
	vars := mux.Vars(r)

	link, err := a.service.Link(context.Background(), vars["shortURL"])
	if err != nil {
		resp.Code = http.StatusNotFound
		resp.Error = err.Error()
		api.RenderJSON(w, resp)
		return
	}

	resp.Data = link.Url
	api.RenderJSON(w, resp)
}

func (a *Api) newLink(w http.ResponseWriter, r *http.Request) {
	resp := &api.Response{Code: http.StatusOK}
	params := make(map[string]string)

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		resp.Code = http.StatusUnprocessableEntity
		resp.Error = err.Error()
		api.RenderJSON(w, resp)
		return
	}

	shortLink, err := a.service.NewLink(context.Background(), params["url"])
	if err != nil {
		resp.Code = http.StatusUnprocessableEntity
		resp.Error = err.Error()
		api.RenderJSON(w, resp)
		return
	}

	resp.Data = shortLink
	api.RenderJSON(w, resp)
}

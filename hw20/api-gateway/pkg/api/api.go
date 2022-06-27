package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kevin-glare/hardcode-dev-go/hw20/api-gateway/pkg/service"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/api"
	"log"
	"net/http"
	"time"
)

type Api struct {
	router  *mux.Router
	addr    string
	service *service.ApiService
}

func Run(addr string) {
	api := &Api{
		router: mux.NewRouter(),
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
	a.router.HandleFunc("/{shortURL}", a.redirectToLink).Methods(http.MethodGet)
	a.router.HandleFunc("/api/v1/links", a.addLink).Methods(http.MethodPost)
	a.router.HandleFunc("/api/v1/links/{shortURL}", a.showLink).Methods(http.MethodGet)
	a.router.HandleFunc("/api/v1/statistic", a.statistic).Methods(http.MethodGet)
}

func (a *Api) redirectToLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := a.service.ShowLink(vars["shortURL"])
	if err != nil {
		resp := &api.Response{Code: http.StatusUnprocessableEntity}
		resp.Error = err.Error()
		api.RenderJSON(w, resp)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("%v", result["data"]), http.StatusMovedPermanently)
}

func (a *Api) statistic(w http.ResponseWriter, r *http.Request) {
	resp := &api.Response{Code: http.StatusOK}
	result, err := a.service.Statistic()
	if err != nil {
		resp.Error = err.Error()
		resp.Code = http.StatusUnprocessableEntity
		api.RenderJSON(w, resp)
		return
	}

	resp.Data = result["data"]
	api.RenderJSON(w, resp)
}

func (a *Api) addLink(w http.ResponseWriter, r *http.Request) {
	resp := &api.Response{Code: http.StatusOK}

	params := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		resp.Code = http.StatusUnprocessableEntity
		resp.Error = err.Error()
		api.RenderJSON(w, resp)
		return
	}

	result, err := a.service.AddLink(params["url"])
	if err != nil {
		resp.Error = err.Error()
		resp.Code = http.StatusUnprocessableEntity
		api.RenderJSON(w, resp)
		return
	}

	resp.Data = result["data"]
	api.RenderJSON(w, resp)
}

func (a *Api) showLink(w http.ResponseWriter, r *http.Request) {
	resp := &api.Response{Code: http.StatusOK}
	vars := mux.Vars(r)

	result, err := a.service.ShowLink(vars["shortURL"])
	if err != nil {
		resp.Error = err.Error()
		resp.Code = http.StatusUnprocessableEntity
		api.RenderJSON(w, resp)
		return
	}

	resp.Data = result["data"]
	api.RenderJSON(w, resp)
}

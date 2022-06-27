package api

import (
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
	a.router.HandleFunc("/{id}", a.redirectToLink).Methods(http.MethodGet)
	a.router.HandleFunc("/api/v1/links", a.addLink).Methods(http.MethodPost)
	a.router.HandleFunc("/api/v1/links/{id}", a.showLink).Methods(http.MethodGet)
	a.router.HandleFunc("/api/v1/statistic", a.statistic).Methods(http.MethodGet)
}

func (a *Api) statistic(w http.ResponseWriter, r *http.Request) {
	resp := a.service.Statistic()

	api.RenderJSON(w, resp)
}

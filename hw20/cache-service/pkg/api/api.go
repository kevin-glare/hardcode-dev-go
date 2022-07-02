package api

import (
	"github.com/gorilla/mux"
	"github.com/kevin-glare/hardcode-dev-go/hw20/cache-service/pkg/cache"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/delivery"
	"log"
	"net/http"
	"time"
)

type Api struct {
	router *mux.Router
	addr   string
	cache  *cache.Cache
}

func Run(addr string, cache *cache.Cache) {
	api := &Api{
		router: mux.NewRouter(),
		cache:  cache,
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
	a.router.HandleFunc("/api/v1/links/{shortURL}", a.link).Methods(http.MethodGet)
}

func (a *Api) link(w http.ResponseWriter, r *http.Request) {
	resp := &delivery.Response{Code: http.StatusOK}
	vars := mux.Vars(r)

	link, err := a.cache.Get(vars["shortURL"])
	if err != nil {
		resp.Code = http.StatusUnprocessableEntity
		resp.Error = err.Error()
		delivery.RenderJSON(w, resp)
		return
	}

	resp.Data = link.Url
	delivery.RenderJSON(w, resp)
}

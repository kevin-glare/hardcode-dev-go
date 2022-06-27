package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kevin-glare/hardcode-dev-go/hw20/cache-service/pkg/cache"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/api"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/kfk"
	"log"
	"net/http"
	"time"
)

type Api struct {
	router   *mux.Router
	addr     string
	cache    *cache.Cache
	consumer *kfk.Consumer
}

func Run(addr string, cache *cache.Cache, consumer *kfk.Consumer) {
	api := &Api{
		router:   mux.NewRouter(),
		cache:    cache,
		consumer: consumer,
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
	a.router.HandleFunc("/api/v1/links", a.link).Methods(http.MethodPost)
}

func (a *Api) link(w http.ResponseWriter, r *http.Request) {
	resp := &api.Response{Code: http.StatusOK}
	params := make(map[string]string)

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		resp.Code = http.StatusUnprocessableEntity
		resp.Error = err.Error()
		api.RenderJSON(w, resp)
		return
	}

	link, err := a.cache.Get(params["short_url"])
	if err != nil {
		resp.Code = http.StatusUnprocessableEntity
		resp.Error = err.Error()
		api.RenderJSON(w, resp)
		return
	}

	resp.Data = link.Url
	api.RenderJSON(w, resp)
}

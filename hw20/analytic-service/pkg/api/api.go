package api

import (
	"github.com/gorilla/mux"
	"github.com/kevin-glare/hardcode-dev-go/hw20/analytic-service/pkg/service"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/delivery"
	"log"
	"net/http"
	"time"
)

type Api struct {
	router *mux.Router
	addr   string
	stat   *service.Statistic
}

func Run(addr string, stat *service.Statistic) {
	api := &Api{
		router: mux.NewRouter(),
		stat:   stat,
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
	a.router.HandleFunc("/api/v1/statistic", a.statistic).Methods(http.MethodGet)
}

func (a *Api) statistic(w http.ResponseWriter, r *http.Request) {
	resp := &delivery.Response{Code: http.StatusOK, Data: a.stat.StatisticData()}
	delivery.RenderJSON(w, resp)
}

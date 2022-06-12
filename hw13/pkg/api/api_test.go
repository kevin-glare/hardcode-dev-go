package api

import (
	"github.com/gorilla/mux"
	"os"
	"testing"
)

var api *Api

func TestMain(m *testing.M) {
	api = &Api{
		router: mux.NewRouter(),
		store:  NewStore(),
	}
	api.endpoints()
	os.Exit(m.Run())
}


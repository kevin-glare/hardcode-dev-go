package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testMux *mux.Router

func TestMain(m *testing.M) {
	testMux = mux.NewRouter()
	endpoints(testMux)
	m.Run()
}

func Test_indexHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/index", nil)
	req.Header.Add("Content-Type", "plain/text")

	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	if rr.Body.String() != "index" {
		t.Errorf("ответ неверен: получили %s, а хотели %s", rr.Body.String(), "index")
	}
}

func Test_docsHandler(t *testing.T) {
	// index
	req := httptest.NewRequest(http.MethodGet, "/docs", nil)
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	payload, _ := json.Marshal(repository.Docs)
	if rr.Body.String() != string(payload) {
		t.Errorf("ответ неверен: получили %s, а хотели %s", rr.Body.String(), string(payload))
	}

	// empty search

	query := "empty"
	req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/docs?q=%s", query), nil)
	req.Header.Add("Content-Type", "application/json")

	rr = httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	payload, _ = json.Marshal(repository.Find(query))
	if rr.Body.String() != string(payload) {
		t.Errorf("ответ неверен: получили %s, а хотели %s", rr.Body.String(), string(payload))
	}

	// go search

	query = "go"
	req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/docs?q=%s", query), nil)
	req.Header.Add("Content-Type", "application/json")

	rr = httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	payload, _ = json.Marshal(repository.Find(query))
	if rr.Body.String() != string(payload) {
		t.Errorf("ответ неверен: получили %s, а хотели %s", rr.Body.String(), string(payload))
	}
}

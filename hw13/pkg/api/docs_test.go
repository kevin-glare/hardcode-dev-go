package api

import (
	"bytes"
	"encoding/json"
	"github.com/kevin-glare/hardcode-dev-go/hw5/pkg/crawler"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestAPI_index(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/docs", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	payload, err := json.Marshal(&Response{Data: api.store.Docs})
	if err != nil {
		t.Errorf(err.Error())
	}
	if rr.Body.String() != string(payload) {
		t.Errorf("ответ неверен: получили %s, а хотели %s", rr.Body, payload)
	}
}

func TestAPI_create(t *testing.T) {
	data := crawler.Document{
		URL: "https://example.com",
		Title: "title",
		Body: "body",
	}
	payload, _ := json.Marshal(data)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/docs", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	if len(api.store.Docs) != 3 {
		t.Errorf("ответ неверен: получили %d, а хотели %d", len(api.store.Docs), 3)
	}
}

func TestAPI_read(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/docs/0", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	payload, err := json.Marshal(&Response{Data: api.store.Docs[0]})
	if err != nil {
		t.Errorf(err.Error())
	}
	if rr.Body.String() != string(payload) {
		t.Errorf("ответ неверен: получили %s, а хотели %s", rr.Body, payload)
	}
}

func TestAPI_update(t *testing.T) {
	d := api.store.Docs[0]

	data := crawler.Document{
		ID: d.ID,
		URL: d.URL,
		Title: "title",
		Body: "body",
	}
	payload, _ := json.Marshal(data)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/docs/0", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	payload, err := json.Marshal(&Response{Data: data})
	if err != nil {
		t.Errorf(err.Error())
	}
	if rr.Body.String() != string(payload) {
		t.Errorf("ответ неверен: получили %s, а хотели %s", rr.Body, payload)
	}
}

func TestAPI_destroy(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/docs/0", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	if len(api.store.Docs) != 1 {
		t.Errorf("ответ неверен: получили %d, а хотели %d", len(api.store.Docs), 1)
	}
}

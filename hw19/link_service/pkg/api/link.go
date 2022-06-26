package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	n       = 8
)

func (a *Api) redirectToLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	a.Lock()
	if url, ok := a.data[vars["id"]]; ok {
		log.Println(url)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	}
	a.Unlock()

	resp := &Response{Code: http.StatusNotFound, Error: "URL not found"}
	renderJSON(w, resp)
}

func (a *Api) addLink(w http.ResponseWriter, r *http.Request) {
	resp := &Response{Code: http.StatusOK}

	id := randSeq()

	params := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		resp.Code = http.StatusUnprocessableEntity
		resp.Error = "bad params"
		renderJSON(w, resp)
		return
	}

	url := params["url"]

	a.Lock()
	a.data[id] = url
	a.Unlock()

	err = a.kfk.SendMessage(url)
	if err != nil {
		log.Println(err.Error())
	}

	resp.Data = fmt.Sprintf("https://%s/%s", a.addr, id)
	renderJSON(w, resp)
}

func (a *Api) showLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resp := &Response{Code: http.StatusOK}

	a.Lock()
	if url, ok := a.data[vars["id"]]; ok {
		resp.Data = url
		renderJSON(w, resp)
		return
	}
	a.Unlock()

	resp.Code = http.StatusNotFound
	resp.Error = "URL not found"
	renderJSON(w, resp)
}

func randSeq() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

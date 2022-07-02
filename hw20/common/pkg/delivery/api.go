package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
	Code  int         `json:"-"`
}

// RenderJSON функция обработки структуры Response и передача ответа клиенту
func RenderJSON(w http.ResponseWriter, resp *Response) {
	log.Printf("Response: %+v", resp)

	payload, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.Code)
	w.Write(payload)
}

type ParsingResponse map[string]interface{}

// HttpGet функция отправки GET запроса и обработка ответа
func HttpGet(url string) (ParsingResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := make(ParsingResponse)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	errorMsg := fmt.Sprintf("%s", result["error"])
	if len(errorMsg) > 0 {
		return nil, errors.New(errorMsg)
	}

	return result, err
}

// HttpPost функция отправки POST запроса и обработка ответа
func HttpPost(url string, params url.Values) (ParsingResponse, error) {
	res, err := http.PostForm(url, params)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := make(ParsingResponse)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	errorMsg := fmt.Sprintf("%s", result["error"])
	if len(errorMsg) > 0 {
		return nil, errors.New(errorMsg)
	}

	return result, nil
}

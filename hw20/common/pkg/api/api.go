package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
	Code  int         `json:"-"`
}

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

func Get(url string) (ParsingResponse, error) {
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

func Post(url string, params map[string]interface{}) (ParsingResponse, error) {
	payload := new(bytes.Buffer)
	err := json.NewEncoder(payload).Encode(params)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
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

	return result, nil
}

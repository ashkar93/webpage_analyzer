package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebScraper(t *testing.T) {

	sendRequest("/api/v1/analyze-webpage?url=https://developer.mozilla.org/en-US/docs-234/Web/HTTP/Status", 404, "expected status : 404 got %v", t)
	sendRequest("/api/v1/analyze-webpage?url=https://developer.mozilla.org/en-US/docs/Web/HTTP/Status", 200, "expected status : 200 got %v", t)
	sendRequest("/api/v1/analyze-webpage?url=", 400, "expected status : 400 got %v", t)

}

func sendRequest(url string, status float64, msg string, t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	WebScraper(w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(data), &jsonMap)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if jsonMap["status"].(float64) != status {
		t.Errorf(msg, jsonMap["status"])
	}

}

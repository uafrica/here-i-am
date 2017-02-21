package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func getRequest(t *testing.T) *http.Request {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	return req
}

func testBodyAndContentType(t *testing.T, req *http.Request, body string, ct string) {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != body {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), body)
	}

	if ctype := rr.Header().Get("Content-Type"); ctype != ct {
		t.Errorf("content type header does not match: got %v want %v",
			ctype, "application/json")
	}
}

func TestIndexHandler(t *testing.T) {
	req := getRequest(t)
	testBodyAndContentType(t, req, "Here I am!", "text/plain")
}

func TestIndexHandlerJson(t *testing.T) {
	req := getRequest(t)
	req.Header.Set("Content-Type", "application/json")
	testBodyAndContentType(t, req, `{"status": "Here I am!"}`, "application/json")
}

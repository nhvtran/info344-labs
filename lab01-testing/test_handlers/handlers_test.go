package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestComplimentHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/compliments", nil)
	if err != nil {
		t.Fatal(err)
	}

	name := "steve"
	q := req.URL.Query()
	q.Add("name", name)
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ComplimentHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned the wrong status code: got %d, want %d", status, http.StatusOK)
	}

	expectedPrefix := fmt.Sprintf("%s, you are", name)
	if !strings.HasPrefix(rr.Body.String(), expectedPrefix) {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expectedPrefix)
	}
}

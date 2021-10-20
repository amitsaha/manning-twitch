package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	mux := http.NewServeMux()
	SetupHandlers(mux)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	// 127.0.0.1:87723
	resp, err := http.Get(ts.URL + "/foo")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	responseBody := string(data)
	t.Log(responseBody)
	// io.Copy
}

func TestHealthCheckHandler(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/healtcheck", nil)
	healthCheckHandler(w, req)

	byteBuf := w.Body

	expectedResponse := "ok"
	if byteBuf.String() != expectedResponse {
		t.Fatalf("Expected: %s, Got: %s\n", expectedResponse, byteBuf.String())
	}
}

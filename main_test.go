package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeRouter(t *testing.T) {
	urls := []string{}
	urls = append(urls, "/loopback/")
	urls = append(urls, "/loopback/hello")

	router := makeRouter()
	for _, url := range urls {
		req := httptest.NewRequest("GET", url, nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestSettingMiddleware(t *testing.T) {
	ts := httptest.NewServer(settingMiddleware(getTestHandler()))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

// getTestHandler returns a http.HandlerFunc for testing http middleware
func getTestHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Test Handler")
	}
	return http.HandlerFunc(fn)
}

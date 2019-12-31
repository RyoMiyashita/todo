package loopback

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestLoopback(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/{message}", Loopback)

	req := httptest.NewRequest("GET", "/hello", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	message := new(Meesage)
	byteArray, _ := ioutil.ReadAll(rec.Body)
	if err := json.Unmarshal(byteArray, message); err != nil {
		t.Error(("not message type"))
		return
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "hello", message.Message)
}

package loopback

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Meesage struct {
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

func Loopback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	message := Meesage{vars["message"], time.Now().In(time.UTC)}

	res, _ := json.Marshal(message)

	w.Write(res)
}

package testserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	EmailAddress string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
}

type fooJsonHandler struct{}

func RunJsonServer() {

	mux := http.NewServeMux()
	mux.Handle("/api/myjson", &fooJsonHandler{})

	http.ListenAndServe(":3000", mux)
}

func (f *fooJsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	print(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "json format error", err)
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))
}

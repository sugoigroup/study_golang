package testserver

import (
	"fmt"
	"net/http"
	strutils "../utils"
)

type fooHandler struct{}

func Runserver() {

		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello World")
		})

		mux.Handle("/foo", &fooHandler{})

		http.ListenAndServe(":3000", mux)

}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, strutils.ToUpperCase("Hello Foo!"))
}
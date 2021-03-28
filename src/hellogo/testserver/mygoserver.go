package testserver

import (
	strutils "../utils"
	"fmt"
	"net/http"
)

type fooHandler struct{}

func RunServer() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", mux)

}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, strutils.ToUpperCase("Hello Foo!"))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}

	w.Header().Set("Content-Type", "text/html") // HTML 헤더 설정
	w.Write(htmlMessage(name))

}

func htmlMessage(msg string) []byte {
	// HTML로 웹 페이지 작성
	html := `
	<html>
	<head>
		<title>Hello</title>
	</head>
	<body>
		<div style='text-align:center'><b>Hello ` + msg + `</b></div>
	</body>
	</html>
	`
	return []byte(html)
}

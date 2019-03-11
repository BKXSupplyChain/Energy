package web

import (
	"net/http"
)

func serveFile(path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	}
}

func Serve() {
	http.HandleFunc("/shooter", serveFile("./web/static/shooter.html"))

	http.ListenAndServe(":80", nil)
}

package web

import (
	"net/http"
)

func Serve() {
	http.HandleFunc("/shooter", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/static/shooter.html")
	})

	http.ListenAndServe(":80", nil)
}

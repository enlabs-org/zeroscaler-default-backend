package main

import (
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		code := r.Header.Get("X-Code")
		namespace := r.Header.Get("X-Namespace")
		service := r.Header.Get("X-Service-Name")

		if code == "404" {
			template.Must(template.ParseFiles("errors/404.html")).Execute(w, nil)
		} else if code == "502" {
			template.Must(template.ParseFiles("errors/502.html")).Execute(w, nil)
		} else if code == "503" {
			template.Must(template.ParseFiles("errors/503.html")).Execute(w, nil)
		} else {
			template.Must(template.ParseFiles("errors/default.html")).Execute(w, nil)
		}

		w.Write([]byte("Code: " + code + "\n"))
		w.Write([]byte("Namespace: " + namespace + "\n"))
		w.Write([]byte("Service: " + service + "\n"))
	})
	http.ListenAndServe(":80", nil)
}

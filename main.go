package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("start chat server...")

	hub := newHub()
	go hub.run()

	http.HandleFunc("/", serveHome)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}

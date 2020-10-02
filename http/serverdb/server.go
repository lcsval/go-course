package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/usuarios/", UsuarioHandler)
	log.Println("Listening on :3001...")
	log.Fatal(http.ListenAndServe(":3001", nil))
}

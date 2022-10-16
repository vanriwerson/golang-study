package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) { // análogo às camadas Routes/Controllers
	w.Write([]byte("Olá, mundo!"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Boas-vindas, usuário!"))
}

func main() {
	go fmt.Println("Ouvindo na porta 5000")

	http.HandleFunc("/home", home)

	http.HandleFunc("/users", users)

	log.Fatal(http.ListenAndServe(":5000", nil)) // subindo servidor na porta 5000
}

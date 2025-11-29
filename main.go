package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Pagina principal</h1>")

}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Acerca de nosotros</h1>")

}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/about", handleAbout)

	log.Println("Escuchando en el puerto 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

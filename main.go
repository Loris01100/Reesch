package main

import(
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Définir les routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	//pour démarrer le serveur
	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenue sur le site de Reesch")
}     

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "A propos")
}
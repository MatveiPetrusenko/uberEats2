package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var index = template.Must(template.ParseFiles("./frontend/html/index.html"))
var restaurants = template.Must(template.ParseFiles("./frontend/html/page_restaurant.html"))
var dummy = template.Must(template.ParseFiles("./frontend/html/dummy.html"))

func main() {
	envFile, err := os.Open("./.env")
	if err != nil {
		log.Fatalln(err)
	}
	defer envFile.Close()

	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		envVar := strings.Split(scanner.Text(), "=")
		os.Setenv(envVar[0], envVar[1])
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("port missing")
	}

	assets := http.FileServer(http.Dir("assets"))
	image := http.FileServer(http.Dir("pic"))

	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", assets))
	mux.Handle("/pic/", http.StripPrefix("/pic/", image))

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/restaurants", restaurantsHandler)
	mux.HandleFunc("/dummy", dummyHandler)

	http.ListenAndServe(":"+port, mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	index.Execute(w, nil)
}

func restaurantsHandler(w http.ResponseWriter, r *http.Request) {
	restaurants.Execute(w, nil)
}

func dummyHandler(w http.ResponseWriter, r *http.Request) {
	dummy.Execute(w, nil)
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/LuluBeatson/go-server/pkg/routes"
	"github.com/gorilla/mux"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8080, "port to listen on")
	flag.Parse()
}

func main() {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./static")))
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/form", formHandler).Methods("POST")
	routes.RegisterMoviesRoutes(r)
	routes.RegisterBookStoreRoutes(r)

	fmt.Println("Starting server at port", port)
	fmt.Printf("http://localhost:%v/", port)
	address := fmt.Sprintf(":%v", port)
	if err := http.ListenAndServe(address, r); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello World")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

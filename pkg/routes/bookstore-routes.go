package routes

import (
	"github.com/LuluBeatson/go-server/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooksHandler).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.GetBookByIdHandler).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBookHandler).Methods("POST")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBookHandler).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBookHandler).Methods("DELETE")
}

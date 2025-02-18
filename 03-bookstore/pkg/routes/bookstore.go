package routes

import (
	"github.com/gorilla/mux"
	"github.com/lokesh1jha/bookstore/pkg/controllers"
)

var ResgisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("GET")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods("DELETE")
}

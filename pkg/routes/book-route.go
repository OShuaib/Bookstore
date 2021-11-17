package routes
 
import (
	"github.com/gorilla/mux"
	"github.com/OShuaib/bookstore/pkg/controllers"
)

var RegisterBookStores = func(router *mux.Router){
	router.HandleFunc( "/book/", controllers.CreateBook).Methods("Post")
	router.HandleFunc("/book/", controllers.GetBook).Methods("Get")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("Get")
	router.HandleFunc("/book/{bookId}", controllers.UpadateBook).Methods("Put")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("Delete")
}
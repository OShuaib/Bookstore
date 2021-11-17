package main

import (
	"log"
	"net/http"
	_"github.com/OShuaib/bookstore/pkg/controllers"
	"github.com/OShuaib/bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main(){
	r := mux.NewRouter()

	routes.RegisterBookStores(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost: 100", r))

}


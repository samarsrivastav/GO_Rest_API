package router

import (
	"github.com/gorilla/mux"
	"github.com/samarsrivastav/mongoapi/controller"
)



// this is making the end points of the api ÷
// which api to call for which function 

func Router()  *mux.Router{
	router:=mux.NewRouter()

	router.HandleFunc("/api/movies",controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie",controller.InsertMovies).Methods("POST")
	router.HandleFunc("/api/movie/{id}",controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}",controller.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/deleteall",controller.DeleteMany).Methods("DELETE")
	return router
}
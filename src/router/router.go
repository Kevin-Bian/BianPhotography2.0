package router

import (
	"github.com/Kevin-Bian/BianPhotography2.0/src/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/apiv2", controller.Greet).Methods("GET", "OPTIONS")
	router.HandleFunc("/apiv2/newphoto", controller.CreatePhoto).Methods("POST", "OPTIONS")
	router.HandleFunc("/apiv2/photo/{id}", controller.GetPhoto).Methods("GET", "OPTIONS")
	router.HandleFunc("/apiv2/photo", controller.GetAllPhoto).Methods("GET", "OPTIONS")
	router.HandleFunc("/apiv2/collage/{id}", controller.GetCollage).Methods("GET", "OPTIONS")

	return router
}

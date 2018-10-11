package main

import (
	"fmt"
	"github.com/Kronin-Cloud/estate-inventory/app"
	"github.com/Kronin-Cloud/estate-inventory/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Root).Methods("GET")

	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.ListUsers).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/login", controllers.Logout).Methods("DELETE")
	router.HandleFunc("/metrics", controllers.Metrics).Methods("GET")
	router.HandleFunc("/healthz", controllers.Healthz).Methods("GET")

	router.HandleFunc("/estate", controllers.CreateEstate).Methods("POST")
	router.HandleFunc("/estate", controllers.ListEstates).Methods("Get")
	router.HandleFunc("/estate/{estateId}", controllers.UpdateEstate).Methods("PUT")
	router.HandleFunc("/estate/{estateId}", controllers.GetEstate).Methods("GET")
	router.HandleFunc("/estate/{estateId}", controllers.DeleteEstate).Methods("DELETE")

	router.HandleFunc("/project", controllers.ListProjects).Methods("GET")
	router.HandleFunc("/project", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/project/{projectId}", controllers.GetProject).Methods("GET")
	router.HandleFunc("/project/{projectId}", controllers.UpdateProject).Methods("PUT")
	router.HandleFunc("/project/{projectId}", controllers.DeleteProject).Methods("DELETE")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}

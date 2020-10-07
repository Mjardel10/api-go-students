package main


import (
	"fmt"
	/* "challenge/database" */
	"github.com/gorilla/mux"
	"log"
	"challenge/api"
	"net/http"
	"os"

)

func main(){
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" //localhost
	}
	/* fmt.Printf("Hola Mundo") 
	db := database.GetConnection() */	

	/* var port string = "3000" */
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiRouter.HandleFunc("/insertar-curso", api.Crear_Curso).Methods("POST")
	apiRouter.HandleFunc("/insertar-estudiante", api.Crear_Estudiante).Methods("POST")
	apiRouter.HandleFunc("/cursos", api.Obtener_Cursos).Methods("GET")
	apiRouter.HandleFunc("/estudiantes", api.Obtener_Estudiantes).Methods("GET")
	apiRouter.HandleFunc("/grafica-cursos", api.Obtener_Grafica_Cursos).Methods("GET")

	fmt.Printf("Servidor Corriendo en el puerto 3000")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
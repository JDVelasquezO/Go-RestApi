package main

import (
	"fmt"
	"github.com/gorilla/mux" // Importamos librería para apis
	"log"
	"net/http"
)

// Creamos nuestro objeto de tareas
type task struct {
	id int `json:id`
	name string `json:name`
	content string `json:content`
}

// Arreglo que tendra cada tarea
type allTasks []task

var tasks = allTasks {
	{
		id: 1,
		name: "task1",
		content: "some content",
	},
}

func indexRoute (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bienvenido a la api en Go")
}

func main() {
	// Creamos router en modo estricto
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute) // Indicamos que funcion se indexa con un endopoint
	// Establecemos puesto y enrutador
	log.Fatal(http.ListenAndServe(":3000", router)) // Imprimimos si hubo un error en el server
}

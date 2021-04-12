package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux" // Importamos librería para apis
	"io/ioutil"
	"log"
	"net/http"
)

// Creamos nuestro objeto de tareas
type task struct {
	ID int `json:id`
	Name string `json:name`
	Content string `json:content`
}

// Arreglo que tendra cada tarea
type allTasks []task

// Primera tarea agregada al array
var tasks = allTasks {
	{
		ID: 1,
		Name: "task1",
		Content: "some content",
	},
}

func indexRoute (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bienvenido a la api en Go")
}

func getTasks (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Agregamos cabecera
	json.NewEncoder(w).Encode(tasks)
}

func createTask (w http.ResponseWriter, r *http.Request) {
	var newTask task // preparamos nuestra variable para una nueva tarea
	reqBody, err := ioutil.ReadAll(r.Body) // Recibimos lo pasado por form, ya sea req o err

	if err != nil { // Manejo de errores
		fmt.Fprint(w, "Inserte datos válidos")
	}

	// Si la informacion es correcta, se la asignamos a la nueva tarea
	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1 // generamos un id autoincrementable
	tasks = append(tasks, newTask) // Agregamos a la lista de tareas, la nueva tarea

	w.Header().Set("Content-Type", "application/json") // Agregamos cabecera
	w.WriteHeader(http.StatusCreated) // Agregamos el codigo de estado
	json.NewEncoder(w).Encode(newTask) // Retornamos la tarea al cliente
}

func main() {
	// Creamos router en modo estricto
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute) // Indicamos que funcion se indexa con un endopoint
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	// Establecemos puesto y enrutador
	log.Fatal(http.ListenAndServe(":3000", router)) // Imprimimos si hubo un error en el server
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux" // Importamos librería para apis
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

func getTask (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Aquí obtenemos lo que el usuario envia por la url
	taskID, err := strconv.Atoi(vars["id"]) // Convierte string a int y cacha el error

	// Manejo de errores
	if err != nil {
		fmt.Fprint(w, "Id invalido")
		return
	}

	// Recorrer cada lista de tareas
	for _, task := range tasks {
		if task.ID == taskID { // Validar id
			w.Header().Set("Content-Type", "application/json") // Agregamos cabecera
			json.NewEncoder(w).Encode(task)
		}
	}
}

func delTask (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Aquí obtenemos lo que el usuario envia por la url
	taskID, err := strconv.Atoi(vars["id"]) // Convierte string a int y cacha el error

	// Manejo de errores
	if err != nil {
		fmt.Fprint(w, "Id invalido")
		return
	}

	// Recorrer cada lista de tareas
	for i, task := range tasks {
		if task.ID == taskID { // Validar id
			// Elimina la tarea como una linked list
			// Lo anterior y posterior se usen
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprint(w, "La tarea con el ID %v ha sido removido muy bien", taskID)
		}
	}
}

func updateTask (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Aquí obtenemos lo que el usuario envia por la url
	taskID, err := strconv.Atoi(vars["id"]) // Convierte string a int y cacha el error
	var updatedTask task

	// Manejo de errores
	if err != nil {
		fmt.Fprint(w, "Id invalido")
		return
	}

	reqBody, err2 := ioutil.ReadAll(r.Body) // Obtenemos todos los atributos a actualizar

	if err2 != nil {
		fmt.Fprint(w, "Insertar datos válidos")
	}
	json.Unmarshal(reqBody, &updatedTask)

	// Recorrer cada lista de tareas
	for i, task := range tasks {
		if task.ID == taskID { // Validar id
			if task.ID == taskID {
				tasks = append(tasks[:i], tasks[i+1:]...)
				updatedTask.ID = taskID
				tasks = append(tasks, updatedTask)

				fmt.Fprint(w, "La tarea con el id %v ha sido actualizada", taskID)
			}
		}
	}
}

func main() {
	// Creamos router en modo estricto
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute) // Indicamos que funcion se indexa con un endopoint
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", delTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	// Establecemos puesto y enrutador
	log.Fatal(http.ListenAndServe(":3333", router)) // Imprimimos si hubo un error en el server
}

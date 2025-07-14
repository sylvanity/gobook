package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

// Task represents a single to-do item.
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// taskStore is an in-memory datastore for tasks.
// It includes a mutex to handle concurrent access safely.
type taskStore struct {
	sync.Mutex
	tasks  map[int]Task
	nextID int
}

// apiServer holds dependencies for the API, like the datastore.
type apiServer struct {
	store *taskStore
}

// newApiServer creates a new apiServer with an initialized taskStore.
func newApiServer() *apiServer {
	store := &taskStore{
		tasks:  make(map[int]Task),
		nextID: 1,
	}
	// Add some initial data
	store.tasks[store.nextID] = Task{ID: store.nextID, Title: "Learn Go", Completed: false}
	store.nextID++
	store.tasks[store.nextID] = Task{ID: store.nextID, Title: "Build a REST API", Completed: false}
	store.nextID++
	return &apiServer{store: store}
}

// loggingMiddleware logs incoming HTTP requests.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// getTasksHandler handles GET requests to /tasks.
func (s *apiServer) getTasksHandler(w http.ResponseWriter, r *http.Request) {
	s.store.Lock()
	defer s.store.Unlock()

	tasks := make([]Task, 0, len(s.store.tasks))
	for _, task := range s.store.tasks {
		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// createTaskHandler handles POST requests to /tasks.
func (s *apiServer) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.store.Lock()
	defer s.store.Unlock()

	task.ID = s.store.nextID
	s.store.tasks[task.ID] = task
	s.store.nextID++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// getTaskHandler handles GET requests to /tasks/{id}.
func (s *apiServer) getTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	s.store.Lock()
	defer s.store.Unlock()

	task, ok := s.store.tasks[id]
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// updateTaskHandler handles PUT requests to /tasks/{id}.
func (s *apiServer) updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.store.Lock()
	defer s.store.Unlock()

	if _, ok := s.store.tasks[id]; !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	updatedTask.ID = id
	s.store.tasks[id] = updatedTask

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

// deleteTaskHandler handles DELETE requests to /tasks/{id}.
func (s *apiServer) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	s.store.Lock()
	defer s.store.Unlock()

	if _, ok := s.store.tasks[id]; !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	delete(s.store.tasks, id)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	server := newApiServer()

	r := mux.NewRouter()
	// Route for the collection
	r.HandleFunc("/tasks", server.getTasksHandler).Methods(http.MethodGet)
	r.HandleFunc("/tasks", server.createTaskHandler).Methods(http.MethodPost)
	// Route for a single item
	r.HandleFunc("/tasks/{id}", server.getTaskHandler).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id}", server.updateTaskHandler).Methods(http.MethodPut)
	r.HandleFunc("/tasks/{id}", server.deleteTaskHandler).Methods(http.MethodDelete)

	// Apply middleware
	loggingRouter := loggingMiddleware(r)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", loggingRouter); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

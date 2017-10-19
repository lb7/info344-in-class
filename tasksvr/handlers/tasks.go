package handlers

import (
	"net/http"
)

//TasksHandler handles requests for the /v1/tasks resource
func (ctx *Context) TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//TODO: parse the `completed` query string param as a boolean
		//and call the GetAll() method on the tasks.Store to get the tasks
	case "POST":
		//TODO: decode the request body into a tasks.NewTask
		//and insert it using the tasks.Store
	default:
		http.Error(w, "method must be GET or POST", http.StatusMethodNotAllowed)
		return
	}
}

//SpecificTaskHandler handles requests for the /v1/tasks/...task-id... resource
func (ctx *Context) SpecificTaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PATCH":
		//TODO: decode the request body into a
		//tasks.TaskUpdates struct and pass that
		//to the Update() method on the tasks.Store
	default:
		http.Error(w, "method must be PATCH", http.StatusMethodNotAllowed)
		return
	}
}

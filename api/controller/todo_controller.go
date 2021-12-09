package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/exception"
	"todo/model"
	"todo/service"
)

type todoController struct {
	TodoService service.TodoService
}

func NewTodoController(service service.TodoService) todoController {
	return todoController{TodoService: service}
}

func (controller *todoController) Router(mux *http.ServeMux) {
	mux.HandleFunc("/api/update-todo", controller.Update)
	mux.HandleFunc("/api/create-todo", controller.Create)
	mux.HandleFunc("/api/todo-list", controller.GetAll)
	mux.HandleFunc("/api/get-todo", controller.GetById)
	mux.HandleFunc("/api/delete-todo", controller.Delete)
}

func (controller *todoController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)

		var payload model.TodoRequest

		err := decoder.Decode(&payload)
		exception.ServerError(err, w)

		data := controller.TodoService.Create(payload)

		response := model.ApiResponse{
			Code:   200,
			Status: "OK",
			Data:   data,
		}

		jsonByte, err := json.Marshal(response)
		exception.PanicIfErr(err)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonByte)

	}
}

func (controller *todoController) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		data := controller.TodoService.GetAll()

		response := model.ApiResponse{
			Code:   200,
			Status: "OK",
			Data:   data,
		}

		jsonByte, err := json.Marshal(response)
		exception.ServerError(err, w)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonByte)

	}
}

func (controller *todoController) GetById(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		query := r.URL.Query()
		idString := query.Get("id")

		id, err := strconv.Atoi(idString)
		exception.ServerError(err, w)

		data := controller.TodoService.GetById(id)

		response := model.ApiResponse{
			Code:   200,
			Status: "OK",
			Data:   data,
		}

		jsonByte, err := json.Marshal(response)
		exception.ServerError(err, w)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonByte)

	}
}

func (controller *todoController) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {

		query := r.URL.Query()
		idString := query.Get("id")

		id, err := strconv.Atoi(idString)
		exception.ServerError(err, w)

		controller.TodoService.Delete(id)

		response := model.ApiResponse{
			Code:   200,
			Status: "OK",
			Data:   "success",
		}

		jsonByte, err := json.Marshal(response)
		exception.ServerError(err, w)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonByte)

	}
}

func (controller *todoController) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		decoder := json.NewDecoder(r.Body)

		var payload model.TodoUpdateRequest

		err := decoder.Decode(&payload)
		exception.ServerError(err, w)
		data := controller.TodoService.Update(payload)

		response := model.ApiResponse{
			Code:   200,
			Status: "OK",
			Data:   data,
		}

		jsonByte, err := json.Marshal(response)
		exception.PanicIfErr(err)

		w.Write(jsonByte)

	}
}

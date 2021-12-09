package service

import (
	"fmt"
	"todo/entity"
	"todo/model"
	"todo/repository"
)

type TodoService interface {
	Create(request model.TodoRequest) (response model.TodoResponse)
	GetAll() (response []model.TodoResponse)
	GetById(id int) (response model.TodoResponse)
	Update(request model.TodoUpdateRequest) (response model.TodoResponse)
	Delete(id int)
}

type todoServiceImp struct {
	TodoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) TodoService {
	return &todoServiceImp{
		TodoRepository: todoRepository,
	}
}

func (service *todoServiceImp) Create(request model.TodoRequest) (response model.TodoResponse) {
	todo := entity.Todo{
		Todo: request.Todo,
	}
	result := service.TodoRepository.Create(todo)
	fmt.Println("service", result)
	response = model.TodoResponse(result)
	return response
}

func (service *todoServiceImp) GetAll() (response []model.TodoResponse) {
	result := service.TodoRepository.GetAll()
	for _, todo := range result {
		response = append(response, model.TodoResponse(todo))
	}
	return response
}

func (service *todoServiceImp) GetById(id int) (response model.TodoResponse) {
	result := service.TodoRepository.GetById(id)
	response = model.TodoResponse(result)
	return response
}

func (service *todoServiceImp) Update(request model.TodoUpdateRequest) (response model.TodoResponse) {
	result := service.TodoRepository.Update(entity.Todo(request))
	response = model.TodoResponse(result)
	return response
}

func (service *todoServiceImp) Delete(id int) {
	service.TodoRepository.Delete(id)
}

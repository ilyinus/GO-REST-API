package services

import (
	"github.com/ilyinus/go-rest-api"
	"github.com/ilyinus/go-rest-api/pkg/repositories"
)

type TodoListService struct {
	repo repositories.TodoList
}

func NewTodoListService(repo repositories.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (t *TodoListService) Create(userId int, list rest.TodoList) (int, error) {
	return t.repo.Create(userId, list)
}

func (t *TodoListService) GetAll(userId int) ([]rest.TodoList, error) {
	return t.repo.GetAll(userId)
}

func (t *TodoListService) GetById(userId, id int) (rest.ListsItem, error) {
	return t.repo.GetById(userId, id)
}

func (t *TodoListService) DeleteList(userId, id int) error {
	return t.repo.DeleteList(userId, id)
}

func (t *TodoListService) Update(userId, id int, input rest.UpdateListInput) error {
	return t.repo.Update(userId, id, input)
}

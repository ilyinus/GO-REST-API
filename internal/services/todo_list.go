package services

import (
	"github.com/ilyinus/go-rest-api/internal/core"
	"github.com/ilyinus/go-rest-api/internal/repositories"
)

type TodoListService struct {
	repo repositories.TodoList
}

func NewTodoListService(repo repositories.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (t *TodoListService) Create(userId int, list core.TodoList) (int, error) {
	return t.repo.Create(userId, list)
}

func (t *TodoListService) GetAll(userId int) ([]core.TodoList, error) {
	return t.repo.GetAll(userId)
}

func (t *TodoListService) GetById(userId, id int) (core.ListsItem, error) {
	return t.repo.GetById(userId, id)
}

func (t *TodoListService) DeleteList(userId, id int) error {
	return t.repo.DeleteList(userId, id)
}

func (t *TodoListService) Update(userId, id int, input core.UpdateListInput) error {
	return t.repo.Update(userId, id, input)
}

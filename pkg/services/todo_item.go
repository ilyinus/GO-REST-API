package services

import (
	"github.com/ilyinus/go-rest-api"
	"github.com/ilyinus/go-rest-api/pkg/repositories"
)

type TodoItemService struct {
	repo     repositories.TodoItem
	listRepo repositories.TodoList
}

func NewTodoItemService(repo repositories.TodoItem, listRepo repositories.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (t *TodoItemService) Create(userId, listId int, input rest.TodoItem) (int, error) {
	_, err := t.listRepo.GetById(userId, listId)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, err
	}

	return t.repo.Create(listId, input)
}

func (t *TodoItemService) GetAll(userId, listId int) ([]rest.TodoItem, error) {
	return t.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (rest.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input rest.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}

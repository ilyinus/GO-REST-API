package services

import (
	"github.com/ilyinus/go-rest-api/internal/core"
	"github.com/ilyinus/go-rest-api/internal/repositories"
)

type TodoItemService struct {
	repo     repositories.TodoItem
	listRepo repositories.TodoList
}

func NewTodoItemService(repo repositories.TodoItem, listRepo repositories.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (t *TodoItemService) Create(userId, listId int, input core.TodoItem) (int, error) {
	_, err := t.listRepo.GetById(userId, listId)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, err
	}

	return t.repo.Create(listId, input)
}

func (t *TodoItemService) GetAll(userId, listId int) ([]core.TodoItem, error) {
	return t.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (core.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input core.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}

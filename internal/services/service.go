package services

import (
	"github.com/ilyinus/go-rest-api/internal/core"
	"github.com/ilyinus/go-rest-api/internal/repositories"
)

type Authorization interface {
	CreateUser(user core.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userId int, list core.TodoList) (int, error)
	GetAll(userId int) ([]core.TodoList, error)
	GetById(userId, id int) (core.ListsItem, error)
	DeleteList(userId, id int) error
	Update(userId, id int, input core.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, input core.TodoItem) (int, error)
	GetAll(userId, listId int) ([]core.TodoItem, error)
	GetById(userId, itemId int) (core.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, id int, input core.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
		TodoItem:      NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}

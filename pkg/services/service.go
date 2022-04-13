package services

import (
	"github.com/ilyinus/go-rest-api"
	"github.com/ilyinus/go-rest-api/pkg/repositories"
)

type Authorization interface {
	CreateUser(user rest.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userId int, list rest.TodoList) (int, error)
	GetAll(userId int) ([]rest.TodoList, error)
	GetById(userId, id int) (rest.ListsItem, error)
	DeleteList(userId, id int) error
	Update(userId, id int, input rest.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, input rest.TodoItem) (int, error)
	GetAll(userId, listId int) ([]rest.TodoItem, error)
	GetById(userId, itemId int) (rest.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, id int, input rest.UpdateItemInput) error
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

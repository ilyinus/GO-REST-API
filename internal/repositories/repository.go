package repositories

import (
	"github.com/ilyinus/go-rest-api/internal/core"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user core.User) (int, error)
	GetUser(username, password string) (core.User, error)
}

type TodoList interface {
	Create(userId int, list core.TodoList) (int, error)
	GetAll(userId int) ([]core.TodoList, error)
	GetById(userId, id int) (core.ListsItem, error)
	DeleteList(userId, id int) error
	Update(userId, id int, input core.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, input core.TodoItem) (int, error)
	GetAll(userId, listId int) ([]core.TodoItem, error)
	GetById(userId, itemId int) (core.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, id int, input core.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}

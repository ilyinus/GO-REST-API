package repositories

import (
	"fmt"
	"github.com/ilyinus/go-rest-api"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (t *TodoListPostgres) Create(userId int, list rest.TodoList) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var listId int
	listQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoListsTable)
	row := t.db.QueryRow(listQuery, list.Title, list.Description)
	if err := row.Scan(&listId); err != nil {
		tx.Rollback()
		return 0, err
	}

	userListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) values ($1, $2) RETURNING id", usersListsTable)
	_, err = t.db.Exec(userListQuery, userId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listId, tx.Commit()
}

func (t *TodoListPostgres) GetAll(userId int) ([]rest.TodoList, error) {
	var lists []rest.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)
	err := t.db.Select(&lists, query, userId)

	return lists, err
}

func (t *TodoListPostgres) GetById(userId, id int) (rest.ListsItem, error) {
	var listItem rest.ListsItem

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
								INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		todoListsTable, usersListsTable)
	err := t.db.Get(&listItem, query, userId, id)

	return listItem, err
}

func (t *TodoListPostgres) DeleteList(userId, id int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		todoListsTable, usersListsTable)
	_, err := t.db.Exec(query, userId, id)

	return err
}

func (t *TodoListPostgres) Update(userId, listId int, input rest.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	// title=$1
	// description=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := t.db.Exec(query, args...)
	return err
}

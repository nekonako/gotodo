package repository

import (
	"database/sql"
	"fmt"
	"todo/config"
	"todo/entity"
	"todo/exception"
)

type TodoRepository interface {
	Create(todo entity.Todo) entity.Todo
	GetAll() []entity.Todo
	GetById(id int) entity.Todo
	Update(todo entity.Todo) entity.Todo
	Delete(id int)
}

type todoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{
		DB: db,
	}
}

func (t *todoRepository) Create(todo entity.Todo) entity.Todo {

	ctx, cancel := config.NewDbContext()
	defer cancel()

	fmt.Println(todo)

	var id int
	query := "INSERT INTO todos(todo) VALUES ($1) RETURNING id"
	err := t.DB.QueryRowContext(ctx, query, todo.Todo).Scan(&id)
	exception.PanicIfErr(err)

	todo.Id = id
	return todo
}

func (t *todoRepository) GetAll() []entity.Todo {

	ctx, cancel := config.NewDbContext()
	defer cancel()

	query := "SELECT * FROM todos ORDER BY id DESC"
	rows, err := t.DB.QueryContext(ctx, query)
	exception.PanicIfErr(err)

	defer rows.Close()
	var todos []entity.Todo
	for rows.Next() {
		var todo entity.Todo
		err := rows.Scan(&todo.Id, &todo.Todo)
		exception.PanicIfErr(err)
		todos = append(todos, todo)
	}

	return todos
}

func (t *todoRepository) GetById(id int) entity.Todo {

	ctx, cancel := config.NewDbContext()
	defer cancel()

	query := "SELECT * FROM todos WHERE id =$1"
	rows, err := t.DB.QueryContext(ctx, query, id)
	exception.PanicIfErr(err)

	defer rows.Close()
	var todo entity.Todo
	for rows.Next() {
		err := rows.Scan(&todo.Id, &todo.Todo)
		exception.PanicIfErr(err)
	}
	return todo
}

func (t *todoRepository) Update(todo entity.Todo) entity.Todo {

	ctx, cancel := config.NewDbContext()
	defer cancel()

	query := "UPDATE todos SET todo=$1 WHERE id = $2 RETURNING id, todo"
	rows, err := t.DB.QueryContext(ctx, query, todo.Todo, todo.Id)
	exception.PanicIfErr(err)

	defer rows.Close()
	var newTodo entity.Todo
	for rows.Next() {
		err := rows.Scan(&newTodo.Id, &newTodo.Todo)
		exception.PanicIfErr(err)
	}

	return newTodo
}

func (t *todoRepository) Delete(id int) {

	ctx, cancel := config.NewDbContext()
	defer cancel()

	query := "DELETE FROM todos WHERE id = $1"
	_, err := t.DB.ExecContext(ctx, query, id)
	exception.PanicIfErr(err)

}

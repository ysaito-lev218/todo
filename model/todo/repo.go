package todo

import (
	"github.com/kroton/todo/repo"
	"github.com/russross/meddler"
)

func all(db repo.DB) []*Todo {
	var todo []*Todo
	meddler.QueryAll(db, &todo, "select * from todo")
	return todo
}

func create(db repo.DB, todo *Todo) error {
	return meddler.Insert(db, "todo", todo)
}

func delete(db repo.DB, todo *Todo) error {
	_, err := db.Exec("delete from todo where id = ?", todo.ID)
	return err
}

func finish(db repo.DB, todo *Todo) error {
	_, err := db.Exec("update todo set finished = ? where id = ?", true, todo.ID)
	if err != nil {
		return err
	}

	todo.Finished = true
	return nil
}

func countByFinished(db repo.DB, finished bool) (int, error) {
	var n int
	err := db.QueryRow("select count(*) from todo where finished = ?", finished).Scan(&n)
	return n, err
}
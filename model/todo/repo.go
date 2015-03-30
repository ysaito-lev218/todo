package todo
import (
    "github.com/kroton/todo/repo"
    "github.com/russross/meddler"
)

func All(db repo.DB) []*Todo {
    var todo []*Todo
    meddler.QueryAll(db, &todo, "select * from todo")
    return todo
}

func Create(db repo.DB, todo *Todo) error {
    return meddler.Insert(db, "todo", todo)
}

func Delete(db repo.DB, todo *Todo) error {
    _, err := db.Exec("delete from todo where id = ?", todo.ID)
    return err
}

func Finish(db repo.DB, todo *Todo) error {
    _, err := db.Exec("update todo set finished = ? where id = ?", true, todo.ID)
    if err != nil {
        return err
    }

    todo.Finished = true
    return nil
}

func CountByFinished(db repo.DB, finished bool) (int, error) {
    var n int
    err := db.QueryRow("select count(*) from todo where finished = ?", finished).Scan(&n)
    return n, err
}
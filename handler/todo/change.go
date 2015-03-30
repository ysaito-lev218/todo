package todo

import (
	"strconv"
	"net/http"
	"golang.org/x/net/context"

	"github.com/kroton/todo/repo"
	"github.com/kroton/todo/model/todo"
)

func getID(r *http.Request) (int64, error) {
	return strconv.ParseInt(r.FormValue("id"), 10, 64)
}

func Finish(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if id, err := getID(r); err == nil {
		todo.FinishByID(repo.Con, id)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if id, err := getID(r); err == nil {
		todo.DeleteByID(repo.Con, id)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
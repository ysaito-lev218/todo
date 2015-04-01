package todo

import (
	"net/http"
	"golang.org/x/net/context"

	"github.com/kroton/todo/view"
	"github.com/kroton/todo/model/todo"
)

func Index(ctx context.Context, w http.ResponseWriter, r *http.Request){
	f, ok := ctx.Value(&createFormKey).(createForm)
	if !ok {
		f = createForm{}
	}

	data := struct {
		TodoList []*todo.Todo
		Form     createForm
	}{
		TodoList: todo.All(ctx),
		Form:     f,
	}

	view.Exec(w, "index.html", data)
}
package todo
import (
    "errors"
    "net/http"
    "golang.org/x/net/context"

    "github.com/kroton/todo/repo"
    "github.com/kroton/todo/model/todo"
)

var (
    createFormKey int
)

type createForm struct {
    Title string
    Err   error
}
func (f *createForm) validate() bool {
    n := len(f.Title)
    if n < 1 || 30 < n {
        f.Err = errors.New("タイトルは1文字以上30文字以下")
        return false
    }

    return true
}

func Create(ctx context.Context, w http.ResponseWriter, r *http.Request) {
    f := createForm{ Title: r.FormValue("title"), Err: nil }

    for {
        if !f.validate() {
            break
        }

        if _, err := todo.CreateWithLimit(repo.Con, f.Title, 5); err != nil {
            // 表示用にエラーを修正する
            if lerr, ok := err.(todo.LimitErr); ok {
                f.Err = lerr
            } else {
                f.Err = errors.New("エラーが発生しました")
            }

            break
        }

        // TODO追加に成功した場合はhomeに帰る
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    // 何らかの原因で失敗した場合はフォームをコンテキストに埋め込み、Indexを描画し直す
    Index(context.WithValue(ctx, &createFormKey, f), w, r)
}
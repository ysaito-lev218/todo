package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"html/template"

	"github.com/guregu/db"
	"github.com/guregu/kami"
	"golang.org/x/net/context"

	"github.com/kroton/todo/view"
	"github.com/kroton/todo/handler/todo"
)

func PrePareDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	if _, err := db.Exec("drop table if exists todo"); err != nil {
		panic(err)
	}
	if _, err := db.Exec("create table todo(id integer primary key, title varchar(40), finished boolean)"); err != nil {
		panic(err)
	}

	type init struct {
		Title    string
		Finished bool
	}

	inits := []init {
		init{ Title: "ジョギングする", Finished: false },
		init{ Title: "牛乳を買う", Finished: true},
		init{ Title: "Haskellをやる", Finished: false},
	}

	for _, init := range inits {
		_, err := db.Exec("insert into todo(title, finished) values(?, ?)", init.Title, init.Finished)
		if err != nil {
			panic(err)
		}
	}

	return db
}

func main(){
	dbcon := PrePareDB()
	tmpls := template.Must(template.ParseGlob("./template/*.html"))

	ctx := context.Background()
	ctx = db.WithSQL(ctx, "main", dbcon)
	ctx = view.NewContext(ctx, tmpls)

	// dbを閉じる
	defer db.CloseSQLAll(ctx)

	// 神コンテキスト！
	kami.Context = ctx

	kami.Get("/", todo.Index)
	kami.Post("/create", todo.Create)
	kami.Post("/finish", todo.Finish)
	kami.Post("/delete", todo.Delete)

	kami.Serve()
}
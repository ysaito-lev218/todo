package todo

import (
	"fmt"
	"database/sql"
	"github.com/guregu/db"
	"golang.org/x/net/context"

	"github.com/kroton/todo/repo"
)

// dbコネクションを取り出す
func getDB(ctx context.Context) *sql.DB {
	return db.SQL(ctx, "main")
}

func All(ctx context.Context) []*Todo {
	return all(getDB(ctx))
}

type LimitErr error
func CreateWithLimit(ctx context.Context, title string, limit int) (*Todo, error) {
	var lerr LimitErr = fmt.Errorf("未消化のTODOは%d件しか登録できません", limit)

	// 未消化のTODO件数を調べてlimitを超えないか/超えているかチェックする
	checker := func(db repo.DB, before bool) error {
		n, err := countByFinished(db, false)
		if err != nil {
			return err
		}

		if (before && n >= limit) || (!before && n > limit) {
			return lerr
		}

		return nil
	}

	// DBを取り出す
	mainDB := getDB(ctx)

	// 追加する前にチェックしておく
	if err := checker(mainDB, true); err != nil {
		return nil, err
	}

	// 追加するTODO
	todo := &Todo {
		ID:       0,
		Title:    title,
		Finished: false,
	}

	// タイミングによっては追加したあとlimitを超えるかもしれないのでトランザクションにしておく
	err := repo.Tx(mainDB, func(tx *sql.Tx) error {
		if err := create(tx, todo); err != nil {
			return err
		}
		if err := checker(tx, false); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func FinishByID(ctx context.Context, id int64) error {
	return finish(getDB(ctx), &Todo{ID: id})
}

func DeleteByID(ctx context.Context, id int64) error {
	return delete(getDB(ctx), &Todo{ID: id})
}
package todo

import (
	"fmt"
	"database/sql"

	"github.com/kroton/todo/repo"
)

type LimitErr error

func CreateWithLimit(db repo.DB, title string, limit int) (*Todo, error) {
	var lerr LimitErr = fmt.Errorf("未消化のTODOは%d件しか登録できません", limit)

	// 未消化のTODO件数を調べてlimitを超えないか/超えているかチェックする
	checker := func(db repo.DB, before bool) error {
		n, err := CountByFinished(db, false)
		if err != nil {
			return err
		}

		if (before && n >= limit) || (!before && n > limit) {
			return lerr
		}

		return nil
	}

	// 追加する前にチェックしておく
	if err := checker(db, true); err != nil {
		return nil, err
	}

	// 追加するTODO
	todo := &Todo {
		ID: 0,
		Title: title,
		Finished: false,
	}

	// タイミングによっては追加したあとlimitを超えるかもしれないのでトランザクションにしておく
	err := repo.Tx(db, func(tx *sql.Tx) error {
		if err := Create(tx, todo); err != nil {
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

func FinishByID(db repo.DB, id int64) error {
	return Finish(db, &Todo{ID: id})
}

func DeleteByID(db repo.DB, id int64) error {
	return Delete(db, &Todo{ID: id})
}
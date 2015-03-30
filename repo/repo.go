package repo

import (
	"errors"
	"database/sql"
	"github.com/russross/meddler"
)

// グローバルなDBコネクション
var Con *sql.DB

// QueryRow, Query, Execをまとめたinterface
type DB meddler.DB


func Tx(tdb DB, fn func(*sql.Tx) error) error {
	if tx, ok := tdb.(*sql.Tx); ok {
		// すでにトランザクションに入っている場合は、そのトランザクションに入るのみ
		return fn(tx)
	}

	if db, ok := tdb.(*sql.DB); ok {
		// トランザクション開始
		// fn内でエラーが起こるか Commitに失敗した時にエラーを返す
		tx, err := db.Begin()
		if err != nil {
			return err
		}

		if err := fn(tx); err != nil {
			// fn内でエラーが起こったらRollbackする
			tx.Rollback()
			return err
		}

		if err := tx.Commit(); err != nil {
			// Commitに失敗したらRollbackする (これは必要ないかもしれない)
			tx.Rollback()
			return err
		}

		return nil
	}

	return errors.New("tdb must be *sql.DB or *sql.Tx")
}
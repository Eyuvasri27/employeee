package storages

import (
	"database/sql"
	"goecho/common"
)

type cursor struct {
	Db *sql.DB
}

func Getcursor() *cursor {
	c := new(cursor)
	c.Db = common.Db
	return c
}

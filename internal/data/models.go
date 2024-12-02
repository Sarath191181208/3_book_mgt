package data

import "database/sql"

type Models struct {
	Book BookModel
}

func New(conn *sql.DB) *Models {
	return &Models{
		Book: BookModel{
			conn: conn,
		},
	}
}

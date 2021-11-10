package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type BDMUPPPP struct {
	ID          int64     `json:"id"`
	DateCreated time.Time `json:"date_created"`
	PPPID       int64     `json:"ppp_id"`
}

func (x *BDMUPPPP) InsertTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	date := time.Now()
	return tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s (date_created,ppp_id) VALUES (?,?)", table.BDMUP_PPP), date, x.PPPID)
}

func (x *BDMUPPPP) DeleteTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE ID=?", table.BDMUP_PPP), x.ID)
}

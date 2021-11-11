package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type BDMUPPP struct {
	ID          int64     `json:"id"`
	DateCreated time.Time `json:"date_created"`
	PPPID       int64     `json:"ppp_id"`
}

func (x *BDMUPPP) InsertTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	date := time.Now()

	return tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s(date_created,ppp_id) VALUES (?,?)", table.BDMU_PPP), date, x.PPPID)
}

func (x *BDMUPPP) DeleteTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE ID=?", table.BDMU_PPP), x.ID)
}
package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type KELARP struct {
	ID          int64     `json:"id"`
	DateCreated time.Time `json:"date_created"`
	RPID        int64     `json:"rp_id"`
}

func (x *KELARP) InsertTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	date := time.Now()
	return tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s (date_created,rp_id) VALUES (?,?)", table.KELA_RP), date, x.RPID)
}

func (x *KELARP) DeleteTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id=?", table.KELA_RP), x.ID)
}

package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type PPKRP struct {
	ID          int64     `json:"id"`
	DateCreated time.Time `json:"date_created"`
	RPID        int64     `json:"rp_id"`
}

func (x *PPKRP) InsertTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	date := time.Now()
	return tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s (date_created,rp_id) VALUES (?,?)", table.PPK_RP), date, x.RPID)
}

func (x *PPKRP) DeleteTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id=?", table.PPK_RP), x.ID)
}

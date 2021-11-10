package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type BDMUPRP struct {
	ID          int64     `json:"id"`
	DateCreated time.Time `json:"date_created"`
	RPID        int64     `json:"rp_id"`
}

func (x *BDMUPRP) InsertTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	date := time.Now()
	return tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s (date_created,rp_id) VALUES (?,?)", table.BDMUP_RP), date, x.RPID)
}

func (x *BDMUPRP) DeleteTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id=?", table.BDMUP_RP), x.ID)
}

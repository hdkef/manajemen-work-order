package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type PPKSPK struct {
	ID          int64     `json:"id"`
	DateCreated time.Time `json:"date_created"`
	SPKID       int64     `json:"spk_id"`
}

func (x *PPKSPK) InsertTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	date := time.Now()
	return tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s (date_created,spk_id) VALUES (?,?)", table.PPK_SPK), date, x.SPKID)
}

func (x *PPKSPK) DeleteTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id=?", table.PPK_SPK), x.ID)
}

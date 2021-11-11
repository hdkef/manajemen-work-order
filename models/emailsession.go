package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
)

type EmailSession struct {
	ID    int64 `json:"id"`
	SPKID int64 `json:"spk_id"`
	PIN   int64 `json:"pin"`
}

func (x *EmailSession) InsertTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s (spk_id,pin) VALUES (?,?)", table.EMAIL_SESSION), x.SPKID, x.PIN)
}

func (x *EmailSession) DeleteBySPKIDTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE spk_id=?", table.EMAIL_SESSION), x.SPKID)
}

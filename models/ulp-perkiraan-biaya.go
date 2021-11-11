package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type ULPPerkiraanBiaya struct {
	ID               int64     `json:"id"`
	DateCreated      time.Time `json:"date_created"`
	PerkiraanBiayaID int64     `json:"perkiraan_biaya_id"`
}

func (x *ULPPerkiraanBiaya) InsertTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	date := time.Now()

	return tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s(date_created, perkiraan_biaya_id) VALUES (?,?)", table.ULP_PERKIRAAN_BIAYA), date, x.PerkiraanBiayaID)
}

func (x *ULPPerkiraanBiaya) DeleteTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id=?", table.ULP_PERKIRAAN_BIAYA), x.ID)
}

package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type PPKPengadaan struct {
	ID          int64     `json:"id"`
	DateCreated time.Time `json:"date_created"`
	PengadaanID int64     `json:"pengadaan_id"`
}

func (x *PPKPengadaan) InsertTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	date := time.Now()
	return tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s (date_created,pengadaan_id) VALUES (?,?)", table.PPK_PENGADAAN), date, x.PengadaanID)
}

func (x *PPKPengadaan) DeleteTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id=?", table.PPK_PENGADAAN), x.ID)
}

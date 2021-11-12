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

func (x *PPKPengadaan) Delete(db *sql.DB, ctx context.Context) (sql.Result, error) {
	return db.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE ID=?", table.PPK_PENGADAAN), x.ID)
}

func (x *PPKPengadaan) FindAll(db *sql.DB, ctx context.Context) ([]PPKPengadaan, error) {
	var result []PPKPengadaan
	rows, err := db.QueryContext(ctx, fmt.Sprintf("SELECT id,date_created,pengadaan_id FROM %s", table.PPK_PENGADAAN))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmp PPKPengadaan
		err = rows.Scan(&tmp.ID, &tmp.DateCreated, &tmp.PengadaanID)
		if err != nil {
			return nil, err
		}
		result = append(result, tmp)
	}

	return result, nil
}

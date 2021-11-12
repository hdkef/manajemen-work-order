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
	SPK         SPK       `json:"spk"`
}

func (x *PPKSPK) InsertTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	date := time.Now()
	return tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s (date_created,spk_id) VALUES (?,?)", table.PPK_SPK), date, x.SPKID)
}

func (x *PPKSPK) DeleteTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id=?", table.PPK_SPK), x.ID)
}

func (x *PPKSPK) Delete(db *sql.DB, ctx context.Context) (sql.Result, error) {
	return db.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE ID=?", table.PPK_SPK), x.ID)
}

func (x *PPKSPK) FindAll(db *sql.DB, ctx context.Context) ([]PPKSPK, error) {
	var result []PPKSPK
	rows, err := db.QueryContext(ctx, fmt.Sprintf("SELECT x.id,x.date_created,x.spk_id,y.date_created,y.doc,y.pengadaan_id,y.status,y.worker_email FROM %s AS x JOIN ON %s AS y ON x.spk_id = y.id", table.PPK_SPK, table.SPK))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmp PPKSPK
		err = rows.Scan(&tmp.ID, &tmp.DateCreated, &tmp.SPKID, &tmp.SPK.DateCreated, &tmp.SPK.Doc, &tmp.SPK.PengadaanID, &tmp.SPK.Status, &tmp.SPK.Doc)
		if err != nil {
			return nil, err
		}
		result = append(result, tmp)
	}

	return result, nil
}

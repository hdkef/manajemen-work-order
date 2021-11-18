package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type SPK struct {
	ID          int64     `json:"id"`
	CreatorID   int64     `json:"creator_id"`
	DateCreated time.Time `json:"date_created"`
	Doc         string    `json:"doc"`
	PengadaanID int64     `json:"pengadaan_id"`
	Status      string    `json:"status"`
	WorkerEmail string    `json:"worker_email"`
}

func (x *SPK) InsertTx(tx *sql.Tx, ctx context.Context, creatorid int64) (sql.Result, error) {
	date := time.Now()

	return tx.ExecContext(ctx, fmt.Sprintf("INSERT %s (creator_id,date_created,doc,pengadaan_id,status, worker_email) VALUES (?,?,?,?,?,?)", table.SPK), creatorid, date, x.Doc, x.PengadaanID, x.Status, x.WorkerEmail)
}

func (x *SPK) UpdateStatus(db *sql.DB, ctx context.Context) (sql.Result, error) {
	return db.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET status=? WHERE id=?", table.SPK), x.Status, x.ID)
}

func (x *SPK) UpdateStatusTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET status=? WHERE id=?", table.SPK), x.Status, x.ID)
}

func (x *SPK) FindWorkerEmailTx(tx *sql.Tx, ctx context.Context) (string, error) {
	var workerEmail string

	err := tx.QueryRowContext(ctx, fmt.Sprintf("SELECT worker_email FROM %s WHERE id=?", table.SPK), x.ID).Scan(&workerEmail)
	if err != nil {
		return "", err
	}
	return workerEmail, nil
}

func (x *SPK) FindOne(db *sql.DB, ctx context.Context) (SPK, error) {
	var tmp SPK
	err := db.QueryRowContext(ctx, fmt.Sprintf("SELECT id,creator_id,date_created,doc,pengadaan_id,status,worker_email FROM %s WHERE id=?", table.SPK), x.ID).Scan(&tmp.ID, &tmp.CreatorID, &tmp.DateCreated, &tmp.Doc, &tmp.PengadaanID, &tmp.Status, &tmp.WorkerEmail)
	if err != nil {
		return SPK{}, err
	}
	return tmp, nil
}

func (x *SPK) FindAll(db *sql.DB, ctx context.Context, lastid int64) ([]SPK, error) {
	var result []SPK

	rows, err := db.QueryContext(ctx, fmt.Sprintf("SELECT id,creator_id,date_created,doc,pengadaan_id,status,worker_email FROM %s WHERE id > ? LIMIT 10", table.SPK), lastid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmp SPK
		err = rows.Scan(&tmp.ID, &tmp.CreatorID, &tmp.DateCreated, &tmp.Doc, &tmp.PengadaanID, &tmp.Status, &tmp.WorkerEmail)
		if err != nil {
			return nil, err
		}
		result = append(result, tmp)
	}

	return result, nil
}

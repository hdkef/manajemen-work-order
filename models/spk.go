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

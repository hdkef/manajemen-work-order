package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type Pengadaan struct {
	ID               int64     `json:"id"`
	CreatorID        int64     `json:"creator_id"`
	DateCreated      time.Time `json:"date_created"`
	Doc              string    `json:"doc"`
	PerkiraanBiayaID int64     `json:"perkiraan_biaya_id"`
}

func (x *Pengadaan) InsertTx(tx *sql.Tx, ctx context.Context, creatorid int64) (sql.Result, error) {
	date := time.Now()

	return tx.ExecContext(ctx, fmt.Sprintf("INSERT %s (creator_id,date_created,doc,perkiraan_biaya_id) VALUES (?,?,?,?)", table.PENGADAAN), creatorid, date, x.Doc, x.PerkiraanBiayaID)
}
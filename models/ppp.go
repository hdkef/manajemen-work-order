package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type PPP struct {
	ID          int64     `json:"id"`
	DateCreated time.Time `json:"date_created"`
	CreatorID   int64     `json:"creator_id"`
	Doc         string    `json:"doc"`
	Status      string    `json:"status"`
	Perihal     string    `json:"perihal"`
	Nota        string    `json:"nota"`
	Pekerjaan   string    `json:"pekerjaan"`
	Sifat       string    `json:"sifat"`
	BDMUID      int64     `json:"bdmu_id"`
	BDMUPID     int64     `json:"bdmup_id"`
	KELAID      int64     `json:"kela_id"`
}

func (x *PPP) InsertTx(tx *sql.Tx, ctx context.Context, creatorid int64) (sql.Result, error) {

	date := time.Now()

	return tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s(date_created,creator_id, doc, status, perihal, nota, sifat, pekerjaan) VALUES (?,?,?,?,?,?,?,?)", table.PPP), date, creatorid, x.Doc, x.Status, x.Perihal, x.Nota, x.Sifat, x.Pekerjaan)
}

func (x *PPP) UpdateStatusAndBDMUIDTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET status=?,bdmu_id=? WHERE id=?", table.PPP), x.Status, x.BDMUID, x.ID)
}

func (x *PPP) UpdateStatusAndBDMUPIDTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET status=?,bdmup_id=? WHERE id=?", table.PPP), x.Status, x.BDMUPID, x.ID)
}

func (x *PPP) UpdateStatusAndKELAIDTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET status=?,kela_id=? WHERE id=?", table.PPP), x.Status, x.KELAID, x.ID)
}

func (x *PPP) UpdateStatusTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET status=?", table.PPP), x.Status)
}

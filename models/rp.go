package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type RP struct {
	ID          int64     `json:"id"`
	CreatorID   int64     `json:"creator_id"`
	DateCreated time.Time `json:"date_created"`
	Doc         string    `json:"doc"`
	Status      string    `json:"status"`
	PPPID       int64     `json:"ppp_id"`
	BDMUID      int64     `json:"bdmu_id"`
	BDMUPID     int64     `json:"bdmup_id"`
	KELAID      int64     `json:"kela_id"`
}

func (x *RP) InsertTx(tx *sql.Tx, ctx context.Context, creatorid int64) (sql.Result, error) {
	date := time.Now()

	return tx.ExecContext(ctx, fmt.Sprintf("INSERT %s (creator_id,date_created,doc,status,ppp_id) VALUES (?,?,?,?,?)", table.RP), creatorid, date, x.Doc, x.Status, x.PPPID)
}

func (x *RP) UpdateStatusAndBDMUIDTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET status=?,bdmu_id=? WHERE id=?", table.RP), x.Status, x.BDMUID, x.ID)
}

func (x *RP) UpdateStatusAndBDMUPIDTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET status=?,bdmup_id=? WHERE id=?", table.RP), x.Status, x.BDMUPID, x.ID)
}

func (x *RP) UpdateStatusAndKELAIDTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET status=?,kela_id=? WHERE id=?", table.RP), x.Status, x.KELAID, x.ID)
}

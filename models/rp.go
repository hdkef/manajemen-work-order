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

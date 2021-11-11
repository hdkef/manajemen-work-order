package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type RPRejected struct {
	ID          int64     `json:"id"`
	DateCreated time.Time `json:"date_created"`
	CreatorID   int64     `json:"creator_id"`
	RPID        int64     `json:"rp_id"`
	MSG         string    `json:"msg"`
}

func (x *RPRejected) InsertTx(tx *sql.Tx, ctx context.Context, creatorid int64) (sql.Result, error) {
	date := time.Now()

	return tx.ExecContext(ctx, fmt.Sprintf("INSERT %s (creator_id,date_created,rp_id,msg) VALUES (?,?,?,?)", table.RP_REJECTED), creatorid, date, x.RPID, x.MSG)
}

package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type PerkiraanBiaya struct {
	ID          int64     `json:"id"`
	DateCreated time.Time `json:"date_created"`
	RPID        int64     `json:"rp_id"`
	CreatorID   int64     `json:"creator_id"`
	EstCost     float64   `json:"est_cost"`
	Doc         string    `json:"doc"`
}

func (x *PerkiraanBiaya) InsertTx(tx *sql.Tx, ctx context.Context, creatorid int64) (sql.Result, error) {
	date := time.Now()

	return tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s(date_created, rp_id, est_cost, creator_id, doc) VALUES (?,?,?,?,?)", table.PERKIRAAN_BIAYA), date, x.RPID, x.EstCost, x.CreatorID, x.Doc)
}

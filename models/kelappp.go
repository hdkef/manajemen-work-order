package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type KELAPPP struct {
	ID          int64     `json:"id"`
	DateCreated time.Time `json:"date_created"`
	PPPID       int64     `json:"ppp_id"`
}

func (x *KELAPPP) InsertTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	date := time.Now()
	return tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s (date_created,ppp_id) VALUES (?,?)", table.KELA_PPP), date, x.PPPID)
}

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

func (x *PerkiraanBiaya) FindOne(db *sql.DB, ctx context.Context) (PerkiraanBiaya, error) {
	var tmp PerkiraanBiaya
	err := db.QueryRowContext(ctx, fmt.Sprintf("SELECT id,date_created,rp_id,creator_id,est_cost,doc FROM %s WHERE id=?", table.PERKIRAAN_BIAYA), x.ID).Scan(&tmp.ID, &tmp.DateCreated, &tmp.RPID, &tmp.CreatorID, &tmp.EstCost, &tmp.Doc)
	if err != nil {
		return PerkiraanBiaya{}, err
	}
	return tmp, nil
}

func (x *PerkiraanBiaya) FindAll(db *sql.DB, ctx context.Context, lastid int64) ([]PerkiraanBiaya, error) {
	var result []PerkiraanBiaya

	rows, err := db.QueryContext(ctx, fmt.Sprintf("SELECT id,date_created,rp_id,creator_id,est_cost,doc FROM %s WHERE id > ? LIMIT 10", table.PERKIRAAN_BIAYA), lastid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmp PerkiraanBiaya
		err = rows.Scan(&tmp.ID, &tmp.DateCreated, &tmp.RPID, &tmp.CreatorID, &tmp.EstCost, &tmp.Doc)
		if err != nil {
			return nil, err
		}
		result = append(result, tmp)
	}

	return result, nil
}

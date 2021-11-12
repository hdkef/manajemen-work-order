package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type ULPPerkiraanBiaya struct {
	ID               int64          `json:"id"`
	DateCreated      time.Time      `json:"date_created"`
	PerkiraanBiayaID int64          `json:"perkiraan_biaya_id"`
	PerkiraanBiaya   PerkiraanBiaya `json:"perkiraan_biaya"`
}

func (x *ULPPerkiraanBiaya) InsertTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	date := time.Now()

	return tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s(date_created, perkiraan_biaya_id) VALUES (?,?)", table.ULP_PERKIRAAN_BIAYA), date, x.PerkiraanBiayaID)
}

func (x *ULPPerkiraanBiaya) DeleteTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id=?", table.ULP_PERKIRAAN_BIAYA), x.ID)
}

func (x *ULPPerkiraanBiaya) Delete(db *sql.DB, ctx context.Context) (sql.Result, error) {
	return db.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE ID=?", table.ULP_PERKIRAAN_BIAYA), x.ID)
}

func (x *ULPPerkiraanBiaya) FindAll(db *sql.DB, ctx context.Context) ([]ULPPerkiraanBiaya, error) {
	var result []ULPPerkiraanBiaya
	rows, err := db.QueryContext(ctx, fmt.Sprintf("SELECT x.id,x.date_created,x.perkiraan_biaya_id,y.date_created,y.rp_id,y.creator_id,y.est_cost,y.doc FROM %s AS x JOIN %s AS y ON x.perkiraan_biaya_id = y.id", table.ULP_PERKIRAAN_BIAYA, table.PERKIRAAN_BIAYA))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmp ULPPerkiraanBiaya
		err = rows.Scan(&tmp.ID, &tmp.DateCreated, &tmp.PerkiraanBiayaID, &tmp.PerkiraanBiaya.DateCreated, &tmp.PerkiraanBiaya.RPID, &tmp.PerkiraanBiaya.CreatorID, &tmp.PerkiraanBiaya.EstCost, &tmp.PerkiraanBiaya.Doc)
		if err != nil {
			return nil, err
		}
		result = append(result, tmp)
	}

	return result, nil
}

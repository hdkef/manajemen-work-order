package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type PPKRP struct {
	ID          int64     `json:"id"`
	DateCreated time.Time `json:"date_created"`
	RPID        int64     `json:"rp_id"`
	RP          RP        `json:"rp"`
}

type PPKRPRepo struct {
	ID          int64
	DateCreated time.Time
	RPID        int64
	RP          RPRepo
}

func (x *PPKRP) InsertTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	date := time.Now()
	return tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s (date_created,rp_id) VALUES (?,?)", table.PPK_RP), date, x.RPID)
}

func (x *PPKRP) DeleteTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id=?", table.PPK_RP), x.ID)
}

func (x *PPKRP) Delete(db *sql.DB, ctx context.Context) (sql.Result, error) {
	return db.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE ID=?", table.PPK_RP), x.ID)
}

func (x *PPKRP) Get(db *sql.DB, ctx context.Context) ([]PPKRP, error) {
	var result []PPKRP
	rows, err := db.QueryContext(ctx, fmt.Sprintf("SELECT x.id,x.date_created,x.rp_id,y.creator_id,y.date_created,y.doc,y.status,y.reason,y.bdmu_id,y.bdmup_id,y.kela_id FROM %s AS x JOIN %s AS y ON x.rp_id = y.id", table.PPK_RP, table.RP))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmpRepo PPKRPRepo
		err = rows.Scan(&tmpRepo.ID, &tmpRepo.DateCreated, &tmpRepo.RPID, &tmpRepo.RP.CreatorID, &tmpRepo.RP.DateCreated, &tmpRepo.RP.Doc, &tmpRepo.RP.Status, &tmpRepo.RP.Reason, tmpRepo.RP.BDMUID, &tmpRepo.RP.BDMUPID, &tmpRepo.RP.KELAID)
		if err != nil {
			return nil, err
		}
		tmp := PPKRP{
			ID:          tmpRepo.ID,
			DateCreated: tmpRepo.DateCreated,
			RPID:        tmpRepo.RPID,
			RP: RP{
				CreatorID:   tmpRepo.RP.CreatorID,
				DateCreated: tmpRepo.RP.DateCreated,
				Doc:         tmpRepo.RP.Doc,
				Reason:      tmpRepo.RP.Reason.String,
				Status:      tmpRepo.RP.Status,
				BDMUID:      tmpRepo.RP.BDMUID.Int64,
				BDMUPID:     tmpRepo.RP.BDMUPID.Int64,
				KELAID:      tmpRepo.RP.KELAID.Int64,
			},
		}
		result = append(result, tmp)
	}

	return result, nil
}

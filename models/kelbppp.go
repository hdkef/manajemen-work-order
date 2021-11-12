package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type KELBPPP struct {
	ID          int64     `json:"id"`
	DateCreated time.Time `json:"date_created"`
	PPPID       int64     `json:"ppp_id"`
	PPP         PPP       `json:"ppp"`
}

type KELBPPPRepo struct {
	ID          int64
	DateCreated time.Time
	PPPID       int64
	PPP         PPPRepo
}

func (x *KELBPPP) InsertTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	date := time.Now()
	return tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s (date_created,ppp_id) VALUES (?,?)", table.KELB_PPP), date, x.PPPID)
}

func (x *KELBPPP) DeleteTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id=?", table.KELB_PPP), x.ID)
}

func (x *KELBPPP) Delete(db *sql.DB, ctx context.Context) (sql.Result, error) {
	return db.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE ID=?", table.KELB_PPP), x.ID)
}

func (x *KELBPPP) FindAll(db *sql.DB, ctx context.Context) ([]KELBPPP, error) {
	var result []KELBPPP
	rows, err := db.QueryContext(ctx, fmt.Sprintf("SELECT x.id,x.date_created,x.ppp_id,y.date_created,y.creator_id,y.doc,y.status,y.perihal,y.nota,y.pekerjaan,y.sifat,y.reason,y.bdmu_id,y.bdmup_id,y.kela_id FROM %s AS x JOIN %s AS y ON x.ppp_id = y.id", table.KELB_PPP, table.PPP))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmpRepo KELBPPPRepo
		err = rows.Scan(&tmpRepo.ID, &tmpRepo.DateCreated, &tmpRepo.PPPID, &tmpRepo.PPP.DateCreated, &tmpRepo.PPP.CreatorID, &tmpRepo.PPP.Doc, &tmpRepo.PPP.Status, &tmpRepo.PPP.Perihal, &tmpRepo.PPP.Nota, &tmpRepo.PPP.Pekerjaan, &tmpRepo.PPP.Sifat, &tmpRepo.PPP.Reason, &tmpRepo.PPP.BDMUID, &tmpRepo.PPP.BDMUPID, &tmpRepo.PPP.KELAID)
		if err != nil {
			return nil, err
		}
		tmp := KELBPPP{
			ID:          tmpRepo.ID,
			DateCreated: tmpRepo.DateCreated,
			PPPID:       tmpRepo.PPPID,
			PPP: PPP{
				CreatorID:   tmpRepo.PPP.CreatorID,
				DateCreated: tmpRepo.PPP.DateCreated,
				Doc:         tmpRepo.PPP.Doc,
				Status:      tmpRepo.PPP.Status,
				Perihal:     tmpRepo.PPP.Perihal,
				Nota:        tmpRepo.PPP.Nota,
				Sifat:       tmpRepo.PPP.Sifat,
				Pekerjaan:   tmpRepo.PPP.Pekerjaan,
				BDMUID:      tmpRepo.PPP.BDMUID.Int64,
				BDMUPID:     tmpRepo.PPP.BDMUPID.Int64,
				KELAID:      tmpRepo.PPP.KELAID.Int64,
				Reason:      tmpRepo.PPP.Reason.String,
			},
		}
		result = append(result, tmp)
	}

	return result, nil
}

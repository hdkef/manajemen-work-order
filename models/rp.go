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
	Reason      string    `json:"reason"`
}

type RPRepo struct {
	ID          int64
	CreatorID   int64
	DateCreated time.Time
	Doc         string
	Status      string
	PPPID       int64
	BDMUID      sql.NullInt64
	BDMUPID     sql.NullInt64
	KELAID      sql.NullInt64
	Reason      sql.NullString
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

func (x *RP) UpdateStatusTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET status=? WHERE id=?", table.RP), x.Status, x.ID)
}

func (x *RP) UpdateStatusAndReasonTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET status=?,reason=? WHERE id=?", table.RP), x.Status, x.Reason, x.ID)
}

func (x *RP) FindPPPIDTx(tx *sql.Tx, ctx context.Context) (int64, error) {
	var pppid int64

	err := tx.QueryRowContext(ctx, fmt.Sprintf("SELECT ppp_id FROM %s WHERE id=?", table.RP), x.ID).Scan(&pppid)
	if err != nil {
		return 0, err
	}
	return pppid, nil
}

func (x *RP) FindOne(db *sql.DB, ctx context.Context) (RP, error) {
	var tmpRepo RPRepo
	err := db.QueryRowContext(ctx, fmt.Sprintf("SELECT id,creator_id,date_created,doc,status,ppp_id,bdmu_id,bdmup_id,kela_id FROM %s WHERE id=?", table.RP), x.ID).Scan(&tmpRepo.ID, &tmpRepo.CreatorID, &tmpRepo.DateCreated, &tmpRepo.Doc, &tmpRepo.Status, &tmpRepo.PPPID, &tmpRepo.BDMUID, &tmpRepo.BDMUPID, &tmpRepo.KELAID)
	if err != nil {
		return RP{}, err
	}
	return RP{
		ID:          tmpRepo.ID,
		CreatorID:   tmpRepo.CreatorID,
		DateCreated: tmpRepo.DateCreated,
		Doc:         tmpRepo.Doc,
		Status:      tmpRepo.Status,
		PPPID:       tmpRepo.PPPID,
		BDMUID:      tmpRepo.BDMUID.Int64,
		BDMUPID:     tmpRepo.BDMUPID.Int64,
		KELAID:      tmpRepo.KELAID.Int64,
	}, nil
}

func (x *RP) FindAll(db *sql.DB, ctx context.Context) ([]RP, error) {
	var result []RP

	rows, err := db.QueryContext(ctx, fmt.Sprintf("SELECT id,creator_id,date_created,doc,status,ppp_id,bdmu_id,bdmup_id,kela_id FROM %s", table.RP))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmpRepo RPRepo
		err = rows.Scan(&tmpRepo.ID, &tmpRepo.CreatorID, &tmpRepo.DateCreated, &tmpRepo.Doc, &tmpRepo.Status, &tmpRepo.PPPID, &tmpRepo.BDMUID, &tmpRepo.BDMUPID, &tmpRepo.KELAID)
		if err != nil {
			return nil, err
		}
		tmp := RP{
			ID:          tmpRepo.ID,
			CreatorID:   tmpRepo.CreatorID,
			DateCreated: tmpRepo.DateCreated,
			Doc:         tmpRepo.Doc,
			Status:      tmpRepo.Status,
			PPPID:       tmpRepo.PPPID,
			BDMUID:      tmpRepo.BDMUID.Int64,
			BDMUPID:     tmpRepo.BDMUPID.Int64,
			KELAID:      tmpRepo.KELAID.Int64,
		}
		result = append(result, tmp)
	}
	return result, nil
}

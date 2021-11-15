package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
	"time"
)

type PPP struct {
	ID          int64     `json:"id"`
	DateCreated time.Time `json:"date_created"`
	CreatorID   int64     `json:"creator_id"`
	Doc         string    `json:"doc"`
	Status      string    `json:"status"`
	Perihal     string    `json:"perihal"`
	Nota        string    `json:"nota"`
	Pekerjaan   string    `json:"pekerjaan"`
	Sifat       string    `json:"sifat"`
	BDMUID      int64     `json:"bdmu_id"`
	BDMUPID     int64     `json:"bdmup_id"`
	KELAID      int64     `json:"kela_id"`
	Reason      string    `json:"reason"`
}

type PPPRepo struct {
	ID          int64
	DateCreated time.Time
	CreatorID   int64
	Doc         string
	Status      string
	Perihal     string
	Nota        string
	Pekerjaan   string
	Sifat       string
	BDMUID      sql.NullInt64
	BDMUPID     sql.NullInt64
	KELAID      sql.NullInt64
	Reason      sql.NullString
}

func (x *PPP) InsertTx(tx *sql.Tx, ctx context.Context, creatorid int64) (sql.Result, error) {

	date := time.Now()

	return tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s(date_created,creator_id, doc, status, perihal, nota, sifat, pekerjaan) VALUES (?,?,?,?,?,?,?,?)", table.PPP), date, creatorid, x.Doc, x.Status, x.Perihal, x.Nota, x.Sifat, x.Pekerjaan)
}

func (x *PPP) UpdateStatusAndBDMUIDTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET status=?,bdmu_id=? WHERE id=?", table.PPP), x.Status, x.BDMUID, x.ID)
}

func (x *PPP) UpdateStatusAndBDMUPIDTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET status=?,bdmup_id=? WHERE id=?", table.PPP), x.Status, x.BDMUPID, x.ID)
}

func (x *PPP) UpdateStatusAndKELAIDTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET status=?,kela_id=? WHERE id=?", table.PPP), x.Status, x.KELAID, x.ID)
}

func (x *PPP) UpdateStatusTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET status=? WHERE id=?", table.PPP), x.Status, x.ID)
}

func (x *PPP) UpdateStatusAndReasonTx(tx *sql.Tx, ctx context.Context) (sql.Result, error) {
	return tx.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET status=?,reason=? WHERE id=?", table.PPP), x.Status, x.Reason, x.ID)
}

func (x *PPP) FindOne(db *sql.DB, ctx context.Context) (PPP, error) {
	var tmpRepo PPPRepo
	err := db.QueryRowContext(ctx, fmt.Sprintf("SELECT id,date_created,creator_id,doc,status,perihal,nota,pekerjaan,sifat,bdmu_id,bdmup_id,kela_id FROM %s WHERE id=?", table.PPP), x.ID).Scan(&tmpRepo.ID, &tmpRepo.DateCreated, &tmpRepo.CreatorID, &tmpRepo.Doc, &tmpRepo.Status, &tmpRepo.Perihal, &tmpRepo.Nota, &tmpRepo.Pekerjaan, &tmpRepo.Sifat, &tmpRepo.BDMUID, &tmpRepo.BDMUPID, &tmpRepo.KELAID)
	if err != nil {
		return PPP{}, err
	}
	return PPP{
		ID:          tmpRepo.ID,
		DateCreated: tmpRepo.DateCreated,
		CreatorID:   tmpRepo.CreatorID,
		Doc:         tmpRepo.Doc,
		Status:      tmpRepo.Status,
		Perihal:     tmpRepo.Perihal,
		Nota:        tmpRepo.Nota,
		Pekerjaan:   tmpRepo.Pekerjaan,
		Sifat:       tmpRepo.Sifat,
		BDMUID:      tmpRepo.BDMUID.Int64,
		BDMUPID:     tmpRepo.BDMUPID.Int64,
		KELAID:      tmpRepo.KELAID.Int64,
	}, nil
}

func (x *PPP) FindAll(db *sql.DB, ctx context.Context) ([]PPP, error) {
	var result []PPP
	rows, err := db.QueryContext(ctx, fmt.Sprintf("SELECT id,date_created,creator_id,doc,status,perihal,nota,pekerjaan,sifat,bdmu_id,bdmup_id,kela_id FROM %s", table.PPP))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmpRepo PPPRepo
		err = rows.Scan(&tmpRepo.ID, &tmpRepo.DateCreated, &tmpRepo.CreatorID, &tmpRepo.Doc, &tmpRepo.Status, &tmpRepo.Perihal, &tmpRepo.Nota, &tmpRepo.Pekerjaan, &tmpRepo.Sifat, &tmpRepo.BDMUID, &tmpRepo.BDMUPID, &tmpRepo.KELAID)
		if err != nil {
			return nil, err
		}
		tmp := PPP{
			ID:          tmpRepo.ID,
			DateCreated: tmpRepo.DateCreated,
			CreatorID:   tmpRepo.CreatorID,
			Doc:         tmpRepo.Doc,
			Status:      tmpRepo.Status,
			Perihal:     tmpRepo.Perihal,
			Nota:        tmpRepo.Nota,
			Pekerjaan:   tmpRepo.Pekerjaan,
			Sifat:       tmpRepo.Sifat,
			BDMUID:      tmpRepo.BDMUID.Int64,
			BDMUPID:     tmpRepo.BDMUPID.Int64,
			KELAID:      tmpRepo.KELAID.Int64,
		}
		result = append(result, tmp)
	}

	return result, nil
}

package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"manajemen-work-order/table"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DBUSER string
var DBPASS string
var DBHOST string
var DBPORT string
var DBNAME string

func init() {
	_ = godotenv.Load()
	DBUSER = os.Getenv("DBUSER")
	DBPASS = os.Getenv("DBPASS")
	DBHOST = os.Getenv("DBHOST")
	DBPORT = os.Getenv("DBPORT")
	DBNAME = os.Getenv("DBNAME")
}

func DB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUSER, DBPASS, DBHOST, DBPORT, DBNAME)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	for {
		err := db.Ping()
		if err == nil {
			err = initTable(db)
			if err != nil {
				panic(err)
			}
			return db, nil
		}
		fmt.Println("ping db...", err)
		time.Sleep(5000 * time.Millisecond)
	}
}

func initTable(db *sql.DB) error {

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createTableEntity(tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = insertSuperAdmin(tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createTablePPP(tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createTableBDMUPPP(tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createTableBDMUPPPP(tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createTableKELAPPP(tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createTableKELBPPP(tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createTableRP(tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createTableKELARP(tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createTableBDMUPRP(tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createTableBDMURP(tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createTablePPKRP(tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createTablePerkiraanBiaya(tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createTableULPPerkiraanBiaya(tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createTablePPEPerkiraanBiaya(tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createTablePengadaan(tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createTablePPKPengadaan(tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func createTableEntity(tx *sql.Tx, ctx context.Context) error {
	return tx.QueryRowContext(ctx, table.ENTITY_CREATION).Err()
}

func createTablePPP(tx *sql.Tx, ctx context.Context) error {
	return tx.QueryRowContext(ctx, table.PPP_CREATION).Err()
}

func createTableBDMUPPP(tx *sql.Tx, ctx context.Context) error {
	return tx.QueryRowContext(ctx, table.BDMU_PPP_CREATION).Err()
}

func createTableBDMUPPPP(tx *sql.Tx, ctx context.Context) error {
	return tx.QueryRowContext(ctx, table.BDMUP_PPP_CREATION).Err()
}

func createTableKELAPPP(tx *sql.Tx, ctx context.Context) error {
	return tx.QueryRowContext(ctx, table.KELA_PPP_CREATION).Err()
}

func createTableKELBPPP(tx *sql.Tx, ctx context.Context) error {
	return tx.QueryRowContext(ctx, table.KELB_PPP_CREATION).Err()
}

func createTableRP(tx *sql.Tx, ctx context.Context) error {
	return tx.QueryRowContext(ctx, table.RP_CREATION).Err()
}

func createTableKELARP(tx *sql.Tx, ctx context.Context) error {
	return tx.QueryRowContext(ctx, table.KELA_RP_CREATION).Err()
}

func createTableBDMUPRP(tx *sql.Tx, ctx context.Context) error {
	return tx.QueryRowContext(ctx, table.BDMUP_RP_CREATION).Err()
}

func createTableBDMURP(tx *sql.Tx, ctx context.Context) error {
	return tx.QueryRowContext(ctx, table.BDMU_RP_CREATION).Err()
}

func createTablePPKRP(tx *sql.Tx, ctx context.Context) error {
	return tx.QueryRowContext(ctx, table.PPK_RP_CREATION).Err()
}

func createTablePerkiraanBiaya(tx *sql.Tx, ctx context.Context) error {
	return tx.QueryRowContext(ctx, table.PERKIRAAN_BIAYA_CREATION).Err()
}

func createTableULPPerkiraanBiaya(tx *sql.Tx, ctx context.Context) error {
	return tx.QueryRowContext(ctx, table.ULP_PERKIRAAN_BIAYA_CREATION).Err()
}

func createTablePPEPerkiraanBiaya(tx *sql.Tx, ctx context.Context) error {
	return tx.QueryRowContext(ctx, table.PPE_PERKIRAAN_BIAYA_CREATION).Err()
}

func createTablePengadaan(tx *sql.Tx, ctx context.Context) error {
	return tx.QueryRowContext(ctx, table.PENGADAAN_CREATION).Err()
}

func createTablePPKPengadaan(tx *sql.Tx, ctx context.Context) error {
	return tx.QueryRowContext(ctx, table.PPK_PENGADAAN_CREATION).Err()
}

func insertSuperAdmin(tx *sql.Tx, ctx context.Context) error {

	var admpass = "admin"

	pass, err := HashPassword(&admpass)

	var sqlErr *mysql.MySQLError

	if errors.As(err, &sqlErr) && sqlErr.Number != 1062 {
		return err
	}
	err = tx.QueryRowContext(ctx, fmt.Sprintf("INSERT INTO %s (fullname,email,username,password,role,signature) VALUES (?,?,?,?,?,?)", table.ENTITY), "super admin", "example@example.com", "admin", pass, "Super-Admin", "assets/signature/test.png").Err()
	if errors.As(err, &sqlErr) && sqlErr.Number != 1062 {
		return err
	}
	return nil
}

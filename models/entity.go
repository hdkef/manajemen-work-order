package models

import (
	"context"
	"database/sql"
	"fmt"
	"manajemen-work-order/table"
)

type Entity struct {
	ID        int64  `json:"id"`
	Fullname  string `json:"fullname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Signature string `json:"signature"`
}

func (x *Entity) FindOne(db *sql.DB, ctx context.Context, fieldName string, value string) error {
	return db.QueryRowContext(ctx, fmt.Sprintf("SELECT id,fullname,username,password,email,role,signature FROM %s WHERE %s = ?", table.ENTITY, fieldName), value).Scan(&x.ID, &x.Fullname, &x.Username, &x.Password, &x.Email, &x.Role, &x.Signature)
}

func (x *Entity) FindSignature(db *sql.DB, ctx context.Context) (string, error) {
	var signature string
	err := db.QueryRowContext(ctx, fmt.Sprintf("SELECT signature FROM %s WHERE id = ?", table.ENTITY), x.ID).Scan(&signature)
	if err != nil {
		return "", err
	}
	return signature, nil
}

func (x *Entity) Insert(db *sql.DB, ctx context.Context) error {
	return db.QueryRowContext(ctx, fmt.Sprintf("INSERT INTO %s (fullname,username,password,email,role,signature) VALUES (?,?,?,?,?,?)", table.ENTITY), x.Fullname, x.Username, x.Password, x.Email, x.Role, x.Signature).Err()
}

func (x *Entity) Delete(db *sql.DB, ctx context.Context) error {
	return db.QueryRowContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id=?", table.ENTITY), x.ID).Err()
}

func (x *Entity) ChangePWD(db *sql.DB, ctx context.Context) error {
	return db.QueryRowContext(ctx, fmt.Sprintf("UPDATE %s SET password=? WHERE id=?", table.ENTITY), x.Password, x.ID).Err()
}

func (x *Entity) FindAll(db *sql.DB, ctx context.Context) ([]Entity, error) {
	var result []Entity

	rows, err := db.QueryContext(ctx, fmt.Sprintf("SELECT id,fullname,username,password,email,role,signature FROM %s", table.ENTITY))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmp Entity
		err = rows.Scan(&tmp.ID, &tmp.Fullname, &tmp.Username, &tmp.Password, &tmp.Email, &tmp.Role, &tmp.Signature)
		if err != nil {
			return nil, err
		}
		result = append(result, tmp)
	}

	return result, nil
}

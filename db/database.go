package swagger

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DBConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "2020NhatTuan"
	dbName := "kindergarten"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

//SQLQuery ... Exec insert query sql
func SQLQuery(db *sql.DB, query string) (*sql.Tx, *sql.Stmt, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, nil, err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	return tx, stmt, nil
}

//SQLBegin ... Begin transaction sql
func SQLBegin(db *sql.DB) (*sql.Tx, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

//SQLExec ... Exec transaction sql
func SQLExec(tx *sql.Tx, query string) (*sql.Stmt, error) {
	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return stmt, nil
}

package dbabstract

import (
	"Schorl/SchorlPackageManager/utils"
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	ERROR_FILE_NOT_CREATED = "ERROR_FILE_NOT_CREATED"
)

type DBAbstract struct {
	DB *sql.DB
}

func (dbAbs *DBAbstract) Open(path string) error {
	if !utils.IsFileExist(path) {
		return errors.New(ERROR_FILE_NOT_CREATED)
	}

	var e error
	dbAbs.DB, e = sql.Open("sqlite3", path)

	if e != nil {
		return fmt.Errorf("internal error: %v", e)
	}

	return nil
}

func (dbAbs *DBAbstract) Close() {
	if !dbAbs.IsOpen() {
		return
	}

	dbAbs.DB.Close()
}

func (dbAbs *DBAbstract) IsOpen() bool {
	if dbAbs.DB == nil {
		fmt.Println("internal error: can't access database")
		return false
	}

	return true
}

func (dbAbs *DBAbstract) Query(cmd string, args ...any) (*sql.Rows, error) {
	if !dbAbs.IsOpen() {
		return nil, nil
	}

	rows, e := dbAbs.DB.Query(cmd, args...)

	return rows, e
}

func (dbAbs *DBAbstract) Exec(cmd string, args ...any) (sql.Result, error) {
	if !dbAbs.IsOpen() {
		return nil, nil
	}

	res, e := dbAbs.DB.Exec(cmd, args...)

	return res, e
}

func Create(path string, rootStrict bool) {
	var perm int = 0600

	if !rootStrict {
		perm = 0666
	}

	f, e := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.FileMode(perm))
	if e != nil {
		fmt.Println(e)
	}
	f.Close()

	_, e = sql.Open("sqlite3", path)
	fmt.Println(e)
}

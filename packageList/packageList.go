package packagelist

import (
	dbabstract "Schorl/SchorlPackageManager/dbAbstract"
	"Schorl/SchorlPackageManager/global"
	"Schorl/SchorlPackageManager/utils"
	"fmt"
	"os"
)

type PackageList struct {
	db dbabstract.DBAbstract
}

func (pl *PackageList) Open() error {
	e := pl.db.Open(global.AppDataPath + "packageList.db")
	if e != nil {
		return e
	}

	return nil
}

func (pl *PackageList) Close() {
	pl.db.Close()
}

func InitPackageList() error {
	var db dbabstract.DBAbstract
	var packageListPath string = global.AppDataPath + "packageList.db"
	var e error

	// NOTE: should create this file with root permision
	if !utils.IsFileExist(packageListPath) {
		f, e := os.Create(packageListPath)
		if e != nil {
			fmt.Println("error: creating packageList.db failed")
			return e
		}
		f.Close()
	}

	e = db.Open(packageListPath)
	if e != nil {
		fmt.Println("error: creating packageList.db failed")
		return e
	}

	// check if there is a schorlPMInfoTable
	shouldCreateTab := false
	res, e := db.Query("select * from schorlPMInfoTable;")

	if e == nil {
		if !res.Next() {
			shouldCreateTab = true
		}
	} else {
		shouldCreateTab = true
	}

	if shouldCreateTab {
		_, e = db.Exec(global.SchorlPMInfoTableDB)
		if e != nil {
			fmt.Printf("error: %v\n", e)
		}
	}

	defer db.Close()
	return nil
}

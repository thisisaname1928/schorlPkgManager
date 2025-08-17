package main

import (
	"Schorl/SchorlPackageManager/global"
	"Schorl/SchorlPackageManager/utils"
	"errors"
	"fmt"
	"os"
)

func chownRestrictFile(dir string) error {
	e := os.Chown(dir, 1, 1)
	if e != nil {
		return e
	}

	return nil
}

func mkdirRestrict(path string) error {
	if !utils.IsFileExist(path) {
		e := os.Mkdir(path, global.APP_FILE_RESTRICT_PERM)
		if e != nil {
			return e
		}

		return chownRestrictFile(path)
	}

	return nil
}

func selfInstall(dir string) error {
	fmt.Println("Schorl Package Manger will be installed in", dir, "with uid=1")
	// make sure some directory is there
	if !utils.IsFileExist(dir + "/bin/") {
		return errors.New(dir + "/bin/ not found")
	}

	if !utils.IsFileExist(dir + "/app/") {
		return errors.New(dir + "/app/ not found")
	}

	var appUniqueID = "org.schorl.spm@" + global.CurrentVersionString

	// make some important dir
	path := dir + "/app/" + appUniqueID
	e := mkdirRestrict(path)
	if e != nil {
		return e
	}

	path = dir + "/app/" + appUniqueID + "/bin/"
	e = mkdirRestrict(path)
	if e != nil {
		return e
	}

	path = dir + "/app/" + appUniqueID + "/appData/"
	e = mkdirRestrict(path)
	if e != nil {
		return e
	}

	// install binary
	path = dir + "/app/" + appUniqueID + "/bin/spm"
	e = utils.CopyFile(global.ExecutablePath, path)
	if e != nil {
		return e
	}
	e = chownRestrictFile(path)
	if e != nil {
		return e
	}

	// install a link to spm

	return nil
}

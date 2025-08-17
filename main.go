package main

import (
	global "Schorl/SchorlPackageManager/global"
	packagelist "Schorl/SchorlPackageManager/packageList"
	"Schorl/SchorlPackageManager/utils"
	"errors"
	"fmt"
	"os"
	"strings"
)

func createWorkLock() error {
	lockPath := global.AppDataPath + "lock.txt"

	if utils.IsFileExist(lockPath) {
		fmt.Println("error: a lock.txt found, maybe I shouldn't do anythings")
		return errors.New(global.ERROR_PM_LOCKED)
	}

	var e error
	global.LockFile, e = os.Create(lockPath) // should create lock.txt with root only

	return e
}

func deleteWorkLock() {
	os.Remove(global.AppDataPath + "lock.txt")
}

func main() {
	// get some informations
	var e error
	global.ExecutablePath, e = os.Executable()

	if e != nil {
		fmt.Printf("error :%v\n", e)
		return
	}

	tmp := global.ExecutablePath
	textRune := []rune(tmp)
	for textRune[len(textRune)-1] != '/' { // delete file name
		textRune = textRune[:len(textRune)-1]
	}

	tmp = string(textRune) + "../appData/"
	global.AppDataPath = tmp

	// get the first agrument
	i := 1
	for i < len(os.Args) {
		if strings.HasPrefix(os.Args[i], "-") {
			// handle options here
			if i+1 < len(os.Args) {
				switch os.Args[i] {
				case "-setPrefix":
					global.Prefix = os.Args[i+1]
				}
			}
			i += 2
		} else {
			break
		}
	}

	if i >= len(os.Args) {
		fmt.Println("error: read args failed")
		return
	}

	switch os.Args[i] {
	case "test":
		test()
		return
	}

	// delete work lock should be here
	if os.Args[i] == "purge" {
		// do somethings
		deleteWorkLock()
		return
	}

	// handle commands that don't use packageList.db
	switch os.Args[i] {
	case "init":
		fmt.Println("init packageList.db at", global.AppDataPath+"packageList.db")
		e := packagelist.InitPackageList()
		if e != nil {
			fmt.Printf("error: %v\n", e)
		}
		return
	case "selfInstall":
		// need 1 arg, like: spm selfInstall [DIR]
		if i+1 >= len(os.Args) {
			fmt.Println("error: missing dir, u should run somethings like 'spm selfInstall [dir]'")
			return
		}
		e := selfInstall(os.Args[i+1])
		if e != nil {
			fmt.Println("internal error:", e)
		}
		return
	}

	// create work lock
	e = createWorkLock()
	if e != nil {
		return
	}
	defer deleteWorkLock()

	var pkgList packagelist.PackageList
	e = pkgList.Open()
	if e != nil {
		fmt.Println("error: packageList.db not found... should init it first")
		return
	}
	defer pkgList.Close()

}

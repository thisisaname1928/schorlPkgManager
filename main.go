package main

import (
	global "Schorl/SchorlPackageManager/Global"
	dbabstract "Schorl/SchorlPackageManager/dbAbstract"
	"Schorl/SchorlPackageManager/utils"
	"fmt"
	"os"
)

func main() {
	// get some informations
	global.ExecutablePath = utils.GetRealPath(os.Args[0])

	tmp := global.ExecutablePath
	textRune := []rune(tmp)
	for textRune[len(textRune)-1] != '/' { // delete file name
		textRune = textRune[:len(textRune)-1]
	}

	tmp = string(textRune) + "/../appData/"
	global.AppDataPath = tmp

	var db dbabstract.DBAbstract
	e := db.Open("./packageList.db")
	if e != nil {
		fmt.Println(e)
	}
	defer db.Close()
}

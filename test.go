package main

import (
	"Schorl/SchorlPackageManager/utils"
	"fmt"
)

func test() {
	fmt.Println("__TEST__")
	e := utils.CopyFile("test.txt", "test1.txt")
	fmt.Println(e)
}

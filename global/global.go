package global

import (
	"fmt"
	"os"
)

// store global variable

var Prefix string = "/"
var ExecutablePath = ""
var AppDataPath = ""
var CurrentRevision = 1

var LockFile *os.File

const (
	ERROR_PM_LOCKED = "ERROR_PM_LOCKED"
)

var SchorlPMInfoTableDB string = fmt.Sprintf(`create table if not exists schorlPMInfoTable(revision int);
insert into schorlPMInfoTable values (%v);
`, CurrentRevision)

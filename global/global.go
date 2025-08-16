package global

import "fmt"

// store global variable

var Prefix string = "/"
var ExecutablePath = ""
var AppDataPath = ""
var CurrentRevision = 1

var SchorlPMInfoTableDB string = fmt.Sprintf(`create table if not exists schorlPMInfoTable(revision int);
insert into schorlPMInfoTable values (%v);
`, CurrentRevision)

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
var CurrentVersionString = "alpha1"

var LockFile *os.File

const (
	ERROR_PM_LOCKED        = "ERROR_PM_LOCKED"
	APP_FILE_RESTRICT_PERM = 0711
)

// this is unsafe, we should change it

var SchorlPMInfoTableDB string = fmt.Sprintf(`create table if not exists schorlPMInfoTable(revision int);
insert into schorlPMInfoTable values (%v);
`, CurrentRevision)

// create and add it self to packageList.db
var SchorlPackageTableDB string = fmt.Sprintf(`
create table if not exists schorlPackageTable (
	SPMRevision int,
	packageIdentifier text,
	packageUniqueID not null primary key,
	packageName text,
	versionString text,
	versionInt int,
	packageType text
);

insert into schorlPackageTable values (
	%v, 
	"org.schorl.spm",  
	"org.schorl.spm@%v", 
	"Schorl Package Manager", 
	"%v", 
	%v, 
	"sysapp"
);
`, CurrentRevision, CurrentVersionString, CurrentVersionString, CurrentRevision)

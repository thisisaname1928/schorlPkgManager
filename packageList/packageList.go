package packagelist

import dbabstract "Schorl/SchorlPackageManager/dbAbstract"

type PackageList struct {
	db dbabstract.DBAbstract
}

func (pl *PackageList) Open() {

}

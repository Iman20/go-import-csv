package main

import "go-import-csv/go-import-csv/repositories/impl"

const csvBranchFile = ".../center.csv"
const csvUnitFile = ".../unit.csv"
const csvCenterFile = ".../center.csv"

func main() {

	//
	var data = impl.ImportMasterCsv()
	data.ImportCsv(csvBranchFile)
	data.ReadMBranch()

	data = impl.ImportMasterCsv()
	data.ImportCsv(csvUnitFile)
	data.ReadMUnit()

	data = impl.ImportMasterCsv()
	data.ImportCsv(csvCenterFile)
	data.ReadMCenter()
	
}











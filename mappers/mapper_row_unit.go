package mappers

import (
	"log"
	. "go-import-csv/go-import-csv/models"
)

func MapperRowToMUnit(rows [][]string) error {
	var mUnits []MUnit

	for _, fields := range rows {
		mBranch := MUnit{
			BranchCode: fields[2],
			UnitCode:   fields[0],
			UnitName:   fields[1],
			Status:     "A",
		}
		mUnits = append(mUnits, mBranch)
	}

	log.Println(mUnits)

	return nil

}
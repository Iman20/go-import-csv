package mappers

import (
	"log"
	. "go-import-csv/go-import-csv/models"
)


func MapperRowTOMBranch(rows [][]string) error {
	var mBranches []MBranch

	for _, fields := range rows {
		mBranch := MBranch{
			BranchCode: fields[0],
			BranchName: fields[1],
			Status:     "A",
		}
		mBranches = append(mBranches, mBranch)
	}

	log.Println(mBranches)

	return nil

}
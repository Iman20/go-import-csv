package mappers

import (
	"log"
	. "go-import-csv/go-import-csv/models"
)

func MapperRowToMCenter(rows [][]string) error {
	var mCenters []MCenter

	for _, fields := range rows {
		mCenter := MCenter{
			UnitCode:   fields[3],
			CenterCode: fields[0],
			CenterName: fields[1],
			Status:     "A",
		}
		mCenters = append(mCenters, mCenter)
	}

	log.Println(mCenters)

	return nil

}
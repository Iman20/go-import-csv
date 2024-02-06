package impl

import (
	"encoding/csv"
	. "go-import-csv/go-import-csv/mappers"
	. "go-import-csv/go-import-csv/repositories"
	"io"
	"log"
	"os"
)

type importMasterCsvImpl struct {
	rows	[][]string
}

// Init implements IImportMasterCsv.
func (i *importMasterCsvImpl) ImportCsv(path string) error {
	csvReader, _, err := openCsvFile(path)
	if err != nil {
		return err
	}

	rowsData, err := readCsvFile(csvReader)
	if err != nil {
		return err
	}

	i.rows = rowsData
	return nil
}

// ReadMBranch implements IImportMasterCsv.
func (i *importMasterCsvImpl) ReadMBranch() error {
	err := MapperRowTOMBranch(i.rows)
	if err != nil {
		return err
	}
	return nil
}

// ReadMCenter implements IImportMasterCsv.
func (i *importMasterCsvImpl) ReadMCenter() error {
	err := MapperRowToMCenter(i.rows)
	if err != nil {
		return err
	}
	return nil
}

// ReadMUnit implements IImportMasterCsv.
func (i *importMasterCsvImpl) ReadMUnit() error {
	err := MapperRowToMUnit(i.rows)
	if err != nil {
		return err
	}
	return nil
}

func ImportMasterCsv() IImportMasterCsv {
	return &importMasterCsvImpl{}
}


func openCsvFile(path string) (*csv.Reader, *os.File, error) {
	log.Println("=> OPEN CSV FILE")

	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	// new reader with custom delimiter
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	reader.Comma = ';'

	return reader, f, nil
}

func readCsvFile(csvReader *csv.Reader) ([][]string, error) {
	log.Println("=> READ CSV FILE")
	rows, err := csvReader.ReadAll()
	if err != nil {
		if err == io.EOF {
			err = nil
		}
	}
	return rows, nil
}


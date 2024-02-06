package repositories

type IImportMasterCsv interface {
	ImportCsv(path string) error
	ReadMBranch() error
	ReadMUnit() error
	ReadMCenter() error
}

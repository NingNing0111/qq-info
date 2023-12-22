package util

import (
	"encoding/csv"
	"os"

	"github.com/ningning0111/config"
)

var conf = config.GetAppConf()

func OpenCsvFileWriter(f string) *csv.Writer {
	file, err := os.Create(conf.ExportPath + f)
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)

	writer.Flush()
	return writer
}

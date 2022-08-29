package util

import (
	"encoding/csv"
	"fmt"
	"os"
)

func WriteTo(output, path string) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	Handle(err)

	_, err = f.WriteString(output)
	if err != nil {
		Handle(err)
	}
}

func ReadCsv(path string) (error, [][]string) {
	orderInfoFile, err := os.Open(path)
	Handle(err)
	reader := csv.NewReader(orderInfoFile)

	oiRecords, err := reader.ReadAll()
	Handle(err)

	return err, oiRecords
}

func Handle(err error) {
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
}

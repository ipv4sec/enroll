package csv

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"strings"
)

func Read(filename string) ([][]string, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(strings.NewReader(string(dat[:])))

	var records [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		records = append(records, record)
		if err != nil {
			return records, err
		}
	}
	return records, nil
}
package csv

import (
	"encoding/csv"
	"github.com/axgle/mahonia"
	"os"
)

func Read(csvPath string) ([][]string, error) {
	//data, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	return nil, err
	//}
	//r := csv.NewReader(strings.NewReader(string(data[:])))

	//var records [][]string
	//for {
	//	record, err := r.Read()
	//	if err == io.EOF {
	//		break
	//	}
	//	records = append(records, record)
	//	if err != nil {
	//		return records, err
	//	}
	//}
	//return records, nil

	file, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := mahonia.NewDecoder("gbk")
	r := csv.NewReader(decoder.NewReader(file))
	return r.ReadAll()
}
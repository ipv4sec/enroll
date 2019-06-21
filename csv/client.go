package csv

import (
	"encoding/csv"
	"enroll/config"
	"enroll/logger"
	"github.com/axgle/mahonia"
	"github.com/satori/go.uuid"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
	"strings"
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

func Generate(data [][]string) (string, error) {
	csvFilename := uuid.NewV4().String() + ".csv"
	fp, err := os.Create(config.Conf.Csv.Generated + csvFilename)
	if err != nil {
		return csvFilename, err
	}
	defer fp.Close()
	for i:=0; i < len(data); i++ {
		line, err := utf82gdk(strings.Join(data[i], ","))
		if err != nil {
			logger.Error("写入文件时失败:", err.Error())
			continue
		}
		fp.WriteString(line + "\n")
	}
	return csvFilename, nil
}

func utf82gdk(src string) (string, error) {
	reader := transform.NewReader(strings.NewReader(src), simplifiedchinese.GBK.NewEncoder())
	if buf, err := ioutil.ReadAll(reader); err != nil {
		return "", err
	} else {
		return string(buf), nil
	}
}
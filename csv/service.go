package csv

import "errors"

func SaveNameAndAdminId(filename string, adminId int64) error {
	csv := &Csv{
		Name: filename,
		AdminId: adminId}
	rowsAffected, _ := Save(csv)
	if rowsAffected != 0 {
		return nil
	}
	return errors.New("保存上传文件记录失败")
}
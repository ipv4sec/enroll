package user

import (
	"enroll/logger"
	"errors"
)

func SaveImportedCsvDatas(data [][]string, siteId int64) error {
	users := []*User{}
	for i:=0; i < len(data); i++ {
		if len(data[i]) < 4 {
			logger.Info("解析CSV错误, 跳过", data[i])
			continue
		}
		user := User{
			SiteId: siteId,
			Name: data[i][0],
			Num: data[i][1],
			Enroll: data[i][2],
			Major: data[i][3]}
		users = append(users, &user)
	}
	return SaveAll(users)
}

func GetUserBySiteId(siteId int64) ([]*User, error) {
	return FindBySiteId(siteId)
}

func DeleteUserByUserId(userId int64) error {
	rowsAffected := DeleteById(userId)
	if rowsAffected != 1 {
		return errors.New("服务器内部错误")
	}
	return nil
}

func DeleteNotConfirmedUserByUserId(userId int64) error {
	rowsAffected := DeleteByIdAndTag(userId, 1)
	if rowsAffected != 1 {
		return errors.New("服务器内部错误")
	}
	return nil
}

func ConfirmUserBySiteId(siteId int64) int64 {
	return UpdateTagBySiteId(2, siteId)
}
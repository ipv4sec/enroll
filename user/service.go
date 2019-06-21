package user

import (
	"enroll/csv"
	"enroll/logger"
	"enroll/site"
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
	if siteId == 0 {
		return FindAll()
	}
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

func SearchByCardNum(cardNum string) ([]*User, error) {
	return FindByCardNum(cardNum)
}

func CensusByAdmin(admineId int64) ([]*CensusResult, error) {
	//
	if admineId == 1 {
		return FindAllSiteCensus()
	}
	return FindSiteCensusByAdminId(admineId)
}

func GetDownloadCsvFilePathBySiteId(siteId int64) string {
	sites, err := site.FindAll()
	if err != nil {
		logger.Error("查询所有站点出错:", err.Error())
		return ""
	}
	siteMap := map[int64]string{}
	for i:=0; i < len(sites); i++ {
		siteMap[sites[i].Id] = sites[i].Name
	}

	var users []*User
	if siteId == 0 {
		users, err = FindAll()
		if err != nil {
			logger.Error("查询学生失败:", err.Error())
			return ""
		}
	} else {
		users, err = FindBySiteId(siteId)
		if err != nil {
			logger.Error("查询学生失败:", err.Error())
			return ""
		}
	}
	csvData :=[][]string{{"姓名","身份证号","报考层次", "专业名称", "站点"}}
	for i:=0; i < len(users); i++ {
		value, ok :=siteMap[users[i].SiteId]
		if !ok {
			value = ""
		}
		csvData = append(csvData, []string{users[i].Name, users[i].Num, users[i].Enroll, users[i].Major, value})
	}
	csvFilename, err := csv.Generate(csvData)
	if err != nil {
		logger.Error("生成CSV失败:", err.Error())
	}
	return csvFilename
}
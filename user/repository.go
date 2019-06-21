package user

import (
	"enroll/mysql"
	"github.com/t-tiger/gorm-bulk-insert"
)

func FindById(id int64) *User {
	var user User
	mysql.Clinet.First(&user, id)
	if user.Name == "" {
		return nil
	}
	return &user
}

func FindAll() ([]*User, error) {
	var users []*User
	err := mysql.Clinet.Find(&users).Error
	return users, err
}

func SaveAll(users []*User) error {
	records := []interface{}{}
	for i:=0; i < len(users); i++ {
		records = append(records, *users[i])
	}
	return gormbulk.BulkInsert(mysql.Clinet, records, 1000)
}

func FindBySiteId(siteId int64) ([]*User, error) {
	var users []*User
	err := mysql.Clinet.Where("site_id = ?", siteId).Find(&users).Error
	return users, err
}

func DeleteById(id int64) int64 {
	var user User
	user.Id = id
	return mysql.Clinet.Delete(&user).RowsAffected
}

func DeleteByIdAndTag(id int64, tag int64) int64 {
	var user User
	user.Id = id
	return mysql.Clinet.Where("tag = ?", tag).Delete(&user).RowsAffected
}

func UpdateTagBySiteId(tag int64, siteId int64) int64 {
	return mysql.Clinet.Model(User{}).Where("site_id = ?", siteId).Updates(User{Tag: tag}).RowsAffected
}

func FindByCardNum(cardNum string) ([]*User, error) {
	var users []*User
	err := mysql.Clinet.Where("num = ?", cardNum).Find(&users).Error
	return users, err
}

func FindAllSiteCensus() ([]*CensusResult, error) {
	var result []*CensusResult
	mysql.Clinet.
		Raw("SELECT count(users.id) AS num, site_id FROM users GROUP BY site_id").
		Scan(&result)
	return result, nil
}

func FindSiteCensusByAdminId(adminId int64) ([]*CensusResult, error) {
	var result []*CensusResult
	mysql.Clinet.
		Raw("SELECT COUNT(*) as num, site_id FROM users WHERE users.site_id = (SELECT site_id FROM admins WHERE id = ?)", adminId).
		Scan(&result)
	return result, nil
}
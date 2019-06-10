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
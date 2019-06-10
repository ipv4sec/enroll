package site

import (
	"errors"
)

func GetAllSites() ([]*Site, error) {
	return FindAll()
}

func DeleteSiteBySiteId(siteId int64) error {
	rowsAffected := DeleteById(siteId)
	if rowsAffected != 1 {
		return errors.New("服务器内部错误")
	}
	return nil
}

func AddSiteByName(name string) error {
	site := Site{
		Name:name,
	}
	rowsAffected, _ := Save(&site)
	if rowsAffected != 1 {
		return errors.New("服务器内部错误")
	}
	return nil
}
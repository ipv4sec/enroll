package site

import "enroll/mysql"

func FindByName(name string) *Site {
	var site Site
	mysql.Clinet.Where("name = ?", name).First(&site)
	if site.Name == "" {
		return nil
	}
	return &site
}

func FindById(id int64) *Site {
	var site Site
	mysql.Clinet.First(&site, id)
	if site.Name == "" {
		return nil
	}
	return &site
}

// 新增站点
func Save(site *Site) (int64, int64) {
	return mysql.Clinet.Save(&site).RowsAffected, site.Id
}

func FindAll() ([]*Site, error) {
	var sites []*Site
	err := mysql.Clinet.Find(&sites).Error
	return sites, err
}

func DeleteById(id int64) int64 {
	var site Site
	site.Id = id
	return mysql.Clinet.Delete(&site).RowsAffected
}
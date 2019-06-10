package admin

import "enroll/mysql"

func FindById(id int64) *Admin {
	var admin Admin
	mysql.Clinet.First(&admin, id)
	if admin.Name == "" {
		return nil
	}
	return &admin
}

func FindAll() ([]*Admin, error) {
	var admins []*Admin
	err := mysql.Clinet.Find(&admins).Error
	return admins, err
}

func FindByName(name string) *Admin {
	var admin Admin
	mysql.Clinet.Where("name = ?", name).First(&admin)
	if admin.Name == "" {
		return nil
	}
	return &admin
}

func Save(admin *Admin) (int64, int64) {
	return mysql.Clinet.Save(&admin).RowsAffected, admin.Id
}

func DeleteById(id int64) int64 {
	var admin Admin
	admin.Id = id
	return mysql.Clinet.Delete(&admin).RowsAffected
}
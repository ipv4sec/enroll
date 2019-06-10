package admin

import (
	"crypto/md5"
	"encoding/hex"
	"enroll/redis"
	"errors"
	"github.com/gogf/gf/g/util/gvalid"
	"github.com/satori/go.uuid"
)

func GenTokenByNameAndPass(name, pass string) (string, *Admin, error) {
	admin := FindByName(name)
	if admin == nil {
		return "", nil, errors.New("未发现此管理员")
	}
	if admin.Pass != encryptPassword(pass) {
		return "", admin, errors.New("管理员密码错误")
	}
	token := uuid.NewV4().String()
	err := redis.SaveUid(token, admin.Id)
	return token, admin, err
}

func encryptPassword(rawPwd string) string  {
	hasher := md5.New()
	hasher.Write([]byte(rawPwd))
	return hex.EncodeToString(hasher.Sum(nil))
}

func checkPassword(rawPwd string) bool {
	m := gvalid.Check(rawPwd, "password3", nil)
	if m != nil {
		return false
	}
	return true
}


func ChangePasswordByUserId(userId int64, oldPassword, newPassword string) error {
	admin := FindById(userId)
	if admin == nil {
		return errors.New("未发现此管理员")
	}
	if !checkPassword(newPassword) {
		return errors.New("密码格式不合法, 密码格式为任意6-18位的可见字符, 必须包含大小写字母、数字和特殊字符")
	}
	if admin.Pass == encryptPassword(oldPassword) {
		admin.Pass = encryptPassword(newPassword)
		rowsAffected, _ := Save(admin)
		if rowsAffected != 0 {
			return nil
		}
		return errors.New("更改密码失败")
	}
	return errors.New("原密码不匹配")
}

func ChangePasswd(userId int64, newPassword string) error {
	admin := FindById(userId)
	if admin == nil {
		return errors.New("未发现此管理员")
	}
	if !checkPassword(newPassword) {
		return errors.New("密码格式不合法, 密码格式为任意6-18位的可见字符, 必须包含大小写字母、数字和特殊字符")
	}
	admin.Pass = encryptPassword(newPassword)
	rowsAffected, _ := Save(admin)
	if rowsAffected != 0 {
		return nil
	}
	return errors.New("更改密码失败")
}

func GetAllAdmins() ([]*Admin, error) {
	return FindAll()
}

func DeleteAdminByAdminId(adminId int64) error {
	rowsAffected := DeleteById(adminId)
	if rowsAffected != 1 {
		return errors.New("服务器内部错误")
	}
	return nil
}

func AddAdminByNameAndPassAndSiteId(name, pass string, siteId int64) error {
	if !checkPassword(pass) {
		return errors.New("密码格式不合法, 密码格式为任意6-18位的可见字符, 必须包含大小写字母、数字和特殊字符")
	}
	admin := Admin{
		Name:name,
		Pass:encryptPassword(pass),
		SiteId:siteId,
	}
	rowsAffected, _ := Save(&admin)
	if rowsAffected != 1 {
		return errors.New("服务器内部错误")
	}
	return nil
}
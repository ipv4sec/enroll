package admin

type Admin struct {
	Id int64 `json:"id" remark:"主键" gorm:"primary_key;AUTO_INCREMENT"`


	Name string  `json:"name" remark:"姓名" gorm:"UNIQUE_INDEX;NOT NULL"`
	Pass string  `json:"pass" remark:"登录密码" gorm:"NOT NULL"`

	SiteId int64 `json:"siteId" remark:"站点主键" gorm:"DEFAULT:0;NOT NULL"`
}

package user

type User struct {
	Id int64 `json:"id" remark:"主键" gorm:"primary_key;AUTO_INCREMENT"`
	SiteId int64 `json:"siteId" remark:"站点主键" gorm:"NOT NULL"`

	Name string  `json:"name" remark:"姓名" gorm:"NOT NULL"`
	Num string `json:"num" remark:"身份证号" gorm:"UNIQUE_INDEX;NOT NULL"`
	Enroll string `json:"enroll" remark:"报考层次" gorm:"NOT NULL"`
	Major string `json:"major" remark:"专业名称" gorm:"NOT NULL"`

	Tag int64  `json:"tag" remark:"正式状态" gorm:"DEFAULT:1;NOT NULL"` // 2^0 编辑状态 2^1 正式状态
}

type CensusResult struct {
	SiteId int64 `json:"siteId"`
	Num int64 `json:"num"`
}

type DbErr struct {
	Err error
	Data *User
}
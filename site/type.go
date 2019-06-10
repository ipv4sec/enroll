package site

// 站点名称
type Site struct {
	Id int64 `json:"id" remark:"主键" gorm:"primary_key;AUTO_INCREMENT"`
	Name string `json:"name" remark:"名称" gorm:"UNIQUE_INDEX;NOT NULL"`
}

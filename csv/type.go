package csv

type Csv struct {
	Id int64 `json:"id" remark:"主键" gorm:"primary_key;AUTO_INCREMENT"`

	Name string  `json:"name" remark:"上传文件名"`
	AdminId int64 `json:"site_id" remark:"Admin主键"`
}

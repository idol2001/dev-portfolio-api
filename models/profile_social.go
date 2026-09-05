package models

// ProfileSocial 社交链接表
type ProfileSocial struct {
	BaseModel
	Network string `gorm:"column:network;size:256" json:"network"`
	Href    string `gorm:"column:href;size:1024" json:"href"`
	OrderNo int64  `gorm:"column:order_no" json:"order_no"`
}

func (ProfileSocial) TableName() string {
	return "profile_socials"
}

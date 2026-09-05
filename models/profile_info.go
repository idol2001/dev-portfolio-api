package models

// ProfileInfo 个人资料表
type ProfileInfo struct {
	BaseModel
	Name        string `gorm:"column:name;size:256" json:"name"`
	Roles       string `gorm:"column:roles;size:256" json:"roles"`
	About       string `gorm:"column:about;size:4096" json:"about"`
	ImageSource string `gorm:"column:image_source;size:1024" json:"image_source"`
	Logo        string `gorm:"column:logo;size:1024" json:"logo"`
	LogoHeight  int64  `gorm:"column:logo_height" json:"logo_height"`
	LogoWidth   int64  `gorm:"column:logo_width" json:"logo_width"`
	IcpFiling   string `gorm:"column:icp_filing;size:256" json:"icp_filing"`
}

func (ProfileInfo) TableName() string {
	return "profile_infos"
}

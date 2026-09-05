package models

// ProfileNavBar 导航菜单表
type ProfileNavBar struct {
	BaseModel
	Title   string `gorm:"column:title;size:256" json:"title"`
	Href    string `gorm:"column:href;size:256" json:"href"`
	OrderNo int64  `gorm:"column:order_no" json:"order_no"`
}

func (ProfileNavBar) TableName() string {
	return "profile_nav_bars"
}

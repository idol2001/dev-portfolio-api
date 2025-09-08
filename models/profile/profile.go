package profile

import (
	"gorm.io/gorm"
)

type ProfileSkillGroup struct {
	gorm.Model
	SkillGroupTitle string
	SkillGroupItems []ProfileSkill `gorm:"foreignkey:SkillGroupId"`
}

type ProfileSkill struct {
	gorm.Model
	SkillGroupId uint
	SkillIcon    string `gorm:"size:1024"`
	SkillTitle   string
}

type ProfileInfo struct {
	gorm.Model
	Name        string
	Roles       string
	About       string `gorm:"size:4096"`
	ImageSource string `gorm:"size:1024"`
	Logo        string `gorm:"size:1024"`
	LogoHeight  int
	LogoWidth   int
}

type ProfileNavBar struct {
	gorm.Model
	Title   string
	Href    string
	OrderNo int
}

type ProfileSocial struct {
	gorm.Model
	Network string
	Href    string `gorm:"size:1024"`
	OrderNo int
}

var Models = []interface{}{
	&ProfileSkillGroup{},
	&ProfileSkill{},
	&ProfileInfo{},
	&ProfileNavBar{},
	&ProfileSocial{},
}

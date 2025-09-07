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
	SkillIcon    string
	SkillTitle   string
}

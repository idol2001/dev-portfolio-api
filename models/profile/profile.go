package profile

import (
	"dev-portfolio-api/models"
)

// ProfileSkillGroup 技能分类表
type ProfileSkillGroup struct {
	models.BaseModel
	SkillGroupTitle string           `gorm:"column:skill_group_title;size:256" json:"skill_group_title"`
	SkillGroupItems []ProfileSkill   `gorm:"foreignKey:SkillGroupId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"skill_group_items,omitempty"`
}

func (ProfileSkillGroup) TableName() string {
	return "profile_skill_groups"
}

// ProfileSkill 技能项表
type ProfileSkill struct {
	models.BaseModel
	SkillGroupId uint64 `gorm:"column:skill_group_id;index" json:"skill_group_id"`
	SkillIcon    string `gorm:"column:skill_icon;size:256" json:"skill_icon"`
	SkillTitle   string `gorm:"column:skill_title;size:256" json:"skill_title"`
}

func (ProfileSkill) TableName() string {
	return "profile_skills"
}

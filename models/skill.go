package models

// SkillGroup 技能分类
type SkillGroup struct {
	BaseModel
	SkillGroupTitle string       `gorm:"column:skill_group_title;size:256" json:"skill_group_title"`
	SkillGroupItems []SkillItem  `gorm:"foreignKey:SkillGroupId;constraint:OnDelete:SET NULL" json:"skill_group_items,omitempty"`
}

func (SkillGroup) TableName() string {
	return "profile_skill_groups"
}

// SkillItem 技能项
type SkillItem struct {
	BaseModel
	SkillGroupId uint64 `gorm:"column:skill_group_id;index" json:"skill_group_id"`
	SkillIcon    string `gorm:"column:skill_icon;size:256" json:"skill_icon"`
	SkillTitle   string `gorm:"column:skill_title;size:256" json:"skill_title"`
}

func (SkillItem) TableName() string {
	return "profile_skills"
}

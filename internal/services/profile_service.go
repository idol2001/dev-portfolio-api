package services

import (
	profileDto "dev-portfolio-api/internal/dto/profile"
	"dev-portfolio-api/models"
	"dev-portfolio-api/models/profile"
	"dev-portfolio-api/pkg/global"
)

// GetProfile 获取个人资料
func GetProfile() (*models.ProfileInfo, error) {
	profileInfo := &models.ProfileInfo{}
	err := global.DB.First(profileInfo).Error
	if err != nil {
		return nil, err
	}
	return profileInfo, nil
}

// UpdateProfile 更新个人资料
func UpdateProfile(profileInfo *models.ProfileInfo) error {
	return global.DB.Save(profileInfo).Error
}

// GetSocials 获取社交链接列表
func GetSocials() ([]models.ProfileSocial, error) {
	socials := make([]models.ProfileSocial, 0)
	err := global.DB.Order("order_no ASC").Find(&socials).Error
	return socials, err
}

// CreateSocial 创建社交链接
func CreateSocial(social *models.ProfileSocial) error {
	return global.DB.Create(social).Error
}

// UpdateSocial 更新社交链接
func UpdateSocial(social *models.ProfileSocial) error {
	return global.DB.Save(social).Error
}

// DeleteSocial 删除社交链接
func DeleteSocial(id uint) error {
	return global.DB.Delete(&models.ProfileSocial{}, id).Error
}

// GetNavBars 获取导航菜单列表
func GetNavBars() ([]models.ProfileNavBar, error) {
	navBars := make([]models.ProfileNavBar, 0)
	err := global.DB.Order("order_no ASC").Find(&navBars).Error
	return navBars, err
}

// CreateNavBar 创建导航菜单
func CreateNavBar(navBar *models.ProfileNavBar) error {
	return global.DB.Create(navBar).Error
}

// UpdateNavBar 更新导航菜单
func UpdateNavBar(navBar *models.ProfileNavBar) error {
	return global.DB.Save(navBar).Error
}

// DeleteNavBar 删除导航菜单
func DeleteNavBar(id uint) error {
	return global.DB.Delete(&models.ProfileNavBar{}, id).Error
}

// GetSkills 获取技能列表（按分类分组）
func GetSkills() (profileDto.SkillsResp, error) {
	resp := profileDto.SkillsResp{}
	resp.Intro = "I love to learn new things and experiment with new technologies.\nThese are some of the major languages, technologies, tools and platforms I have worked with:"

	list := make([]profile.ProfileSkillGroup, 0)
	query := global.DB
	err := query.Preload("SkillGroupItems").Find(&list).Error
	if err != nil {
		return resp, err
	}
	skills := make([]profileDto.SkillGroup, 0)
	for _, group := range list {
		items := make([]profileDto.SkillItem, 0)
		for _, item := range group.SkillGroupItems {
			items = append(items, profileDto.SkillItem{Icon: item.SkillIcon, Title: item.SkillTitle})
		}
		skills = append(skills, profileDto.SkillGroup{Title: group.SkillGroupTitle, Items: items})
	}
	resp.Skills = skills
	return resp, nil
}

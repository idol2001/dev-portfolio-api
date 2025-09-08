package services

import (
	profileDto "dev-portfolio-api/internal/dto/profile"
	profileDo "dev-portfolio-api/models/profile"
	"dev-portfolio-api/pkg/global"
	"strings"
)

func GetSkills() (profileDto.SkillsResp, error) {
	resp := profileDto.SkillsResp{}
	resp.Intro = "I love to learn new things and experiment with new technologies.\nThese are some of the major languages, technologies, tools and platforms I have worked with:"

	list := make([]profileDo.ProfileSkillGroup, 0)
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

func GetHome() (profileDto.HomeResp, error) {
	resp := profileDto.HomeResp{}
	query := global.DB
	info := profileDo.ProfileInfo{}
	err := query.Model(&profileDo.ProfileInfo{}).First(&info).Error
	if err != nil {
		return resp, err
	}
	resp.Name = info.Name
	resp.Roles = strings.Split(info.Roles, ",")
	return resp, nil
}

func GetAbout() (profileDto.AboutResp, error) {
	resp := profileDto.AboutResp{}
	query := global.DB
	info := profileDo.ProfileInfo{}
	err := query.Model(&profileDo.ProfileInfo{}).First(&info).Error
	if err != nil {
		return resp, err
	}
	resp.About = info.About
	resp.ImageSource = info.ImageSource
	return resp, nil
}

func GetNavBar() (profileDto.NavBarResp, error) {
	resp := profileDto.NavBarResp{}
	info := profileDo.ProfileInfo{}
	query := global.DB
	err := query.Model(&profileDo.ProfileInfo{}).First(&info).Error
	if err != nil {
		return resp, err
	}
	resp.Logo.Source = info.Logo
	resp.Logo.Height = info.LogoHeight
	resp.Logo.Width = info.LogoWidth

	list := make([]profileDo.ProfileNavBar, 0)
	err = query.Model(&profileDo.ProfileNavBar{}).Order("order_no").Find(&list).Error
	if err != nil {
		return resp, err
	}
	items := make([]profileDto.NavItem, 0)
	for _, item := range list {
		items = append(items, profileDto.NavItem{Title: item.Title, Href: item.Href})
	}
	resp.Sections = items
	return resp, nil
}

func GetSocial() (profileDto.SocialResp, error) {
	resp := profileDto.SocialResp{}
	query := global.DB
	list := make([]profileDo.ProfileSocial, 0)
	err := query.Model(&profileDo.ProfileSocial{}).Order("order_no").Find(&list).Error
	if err != nil {
		return resp, err
	}
	items := make([]profileDto.SocialNetwork, 0)
	for _, item := range list {
		items = append(items, profileDto.SocialNetwork{Network: item.Network, Href: item.Href})
	}
	resp.Social = items
	return resp, nil
}

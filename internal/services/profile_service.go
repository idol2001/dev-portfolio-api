package services

import (
	profileDto "dev-portfolio-api/internal/dto/profile"
	profileDo "dev-portfolio-api/models/profile"
	"dev-portfolio-api/pkg/global"
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

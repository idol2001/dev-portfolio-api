package services

import (
	"dev-portfolio-api/models"
	"dev-portfolio-api/pkg/global"
)

// GetProjects 获取所有项目
func GetProjects(publishedOnly bool) ([]models.Project, error) {
	projects := make([]models.Project, 0)
	query := global.DB
	if publishedOnly {
		query = query.Where("status = ?", "published")
	}
	err := query.Order("order_no ASC").Find(&projects).Error
	return projects, err
}

// GetProjectByID 根据 ID 获取项目
func GetProjectByID(id uint) (*models.Project, error) {
	project := &models.Project{}
	err := global.DB.First(project, id).Error
	if err != nil {
		return nil, err
	}
	return project, nil
}

// CreateProject 创建项目
func CreateProject(project *models.Project) error {
	return global.DB.Create(project).Error
}

// UpdateProject 更新项目
func UpdateProject(project *models.Project) error {
	return global.DB.Save(project).Error
}

// DeleteProject 删除项目
func DeleteProject(id uint) error {
	return global.DB.Delete(&models.Project{}, id).Error
}

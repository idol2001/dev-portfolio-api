package handlers

import (
	"dev-portfolio-api/internal/services"
	"dev-portfolio-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetProjects 获取项目列表
func GetProjects(c *gin.Context) {
	published := c.DefaultQuery("published", "true")
	projects, err := services.GetProjects(published == "true")
	if err != nil {
		models.FailWithDetailed("", err.Error(), c)
		return
	}
	models.OkWithData(projects, c)
}

// GetProjectByID 根据 ID 获取项目
func GetProjectByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}
	project, err := services.GetProjectByID(uint(id))
	if err != nil {
		models.FailWithDetailed("", err.Error(), c)
		return
	}
	models.OkWithData(project, c)
}

// CreateProject 创建项目
func CreateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		models.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	if err := services.CreateProject(&project); err != nil {
		models.FailWithMessage("创建失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("创建成功", c)
}

// UpdateProject 更新项目
func UpdateProject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		models.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	project.ID = uint(id)
	if err := services.UpdateProject(&project); err != nil {
		models.FailWithMessage("更新失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("更新成功", c)
}

// DeleteProject 删除项目
func DeleteProject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}
	if err := services.DeleteProject(uint(id)); err != nil {
		models.FailWithMessage("删除失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("删除成功", c)
}

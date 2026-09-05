package handlers

import (
	"dev-portfolio-api/internal/services"
	"dev-portfolio-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetSkillGroups 获取所有技能分类（含子项）
func GetSkillGroups(c *gin.Context) {
	groups, err := services.GetSkillGroups()
	if err != nil {
		models.FailWithMessage("获取失败", c)
		return
	}
	models.OkWithData(groups, c)
}

// CreateSkillGroup 创建技能分类
func CreateSkillGroup(c *gin.Context) {
	var req struct {
		Title string `json:"title" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		models.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	group := &models.SkillGroup{SkillGroupTitle: req.Title}
	if err := services.CreateSkillGroup(group); err != nil {
		models.FailWithMessage("创建失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("创建成功", c)
}

// UpdateSkillGroup 更新技能分类
func UpdateSkillGroup(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}
	var req struct {
		Title string `json:"title" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		models.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	if err := services.UpdateSkillGroup(uint(id), req.Title); err != nil {
		models.FailWithMessage("更新失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("更新成功", c)
}

// DeleteSkillGroup 删除技能分类
func DeleteSkillGroup(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}
	if err := services.DeleteSkillGroup(uint(id)); err != nil {
		models.FailWithMessage("删除失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("删除成功", c)
}

// GetSkillItems 获取指定分类下的技能项
func GetSkillItems(c *gin.Context) {
	groupId, err := strconv.ParseUint(c.Param("groupId"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}
	items, err := services.GetSkillItems(uint(groupId))
	if err != nil {
		models.FailWithMessage("获取失败", c)
		return
	}
	models.OkWithData(items, c)
}

// CreateSkillItem 创建技能项
func CreateSkillItem(c *gin.Context) {
	groupId, err := strconv.ParseUint(c.Param("groupId"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}
	var req struct {
		Icon  string `json:"icon"`
		Title string `json:"title" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		models.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	item := &models.SkillItem{
		SkillGroupId: uint64(groupId),
		SkillIcon:    req.Icon,
		SkillTitle:   req.Title,
	}
	if err := services.CreateSkillItem(item); err != nil {
		models.FailWithMessage("创建失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("创建成功", c)
}

// UpdateSkillItem 更新技能项
func UpdateSkillItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}
	var req struct {
		Icon  string `json:"icon"`
		Title string `json:"title" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		models.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	if err := services.UpdateSkillItem(uint(id), req.Icon, req.Title); err != nil {
		models.FailWithMessage("更新失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("更新成功", c)
}

// DeleteSkillItem 删除技能项
func DeleteSkillItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}
	if err := services.DeleteSkillItem(uint(id)); err != nil {
		models.FailWithMessage("删除失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("删除成功", c)
}

package handlers

import (
	"dev-portfolio-api/internal/services"
	"dev-portfolio-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetProfile 获取个人资料
func GetProfile(c *gin.Context) {
	profileInfo, err := services.GetProfile()
	if err != nil {
		models.FailWithDetailed("", err.Error(), c)
		return
	}
	models.OkWithData(profileInfo, c)
}

// UpdateProfile 更新个人资料
func UpdateProfile(c *gin.Context) {
	var profileInfo models.ProfileInfo
	if err := c.ShouldBindJSON(&profileInfo); err != nil {
		models.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	if err := services.UpdateProfile(&profileInfo); err != nil {
		models.FailWithMessage("更新失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("更新成功", c)
}

// GetSocials 获取社交链接列表
func GetSocials(c *gin.Context) {
	socials, err := services.GetSocials()
	if err != nil {
		models.FailWithDetailed("", err.Error(), c)
		return
	}
	models.OkWithData(socials, c)
}

// CreateSocial 创建社交链接
func CreateSocial(c *gin.Context) {
	var social models.ProfileSocial
	if err := c.ShouldBindJSON(&social); err != nil {
		models.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	if err := services.CreateSocial(&social); err != nil {
		models.FailWithMessage("创建失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("创建成功", c)
}

// UpdateSocial 更新社交链接
func UpdateSocial(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}
	var social models.ProfileSocial
	if err := c.ShouldBindJSON(&social); err != nil {
		models.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	social.ID = uint(id)
	if err := services.UpdateSocial(&social); err != nil {
		models.FailWithMessage("更新失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("更新成功", c)
}

// DeleteSocial 删除社交链接
func DeleteSocial(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}
	if err := services.DeleteSocial(uint(id)); err != nil {
		models.FailWithMessage("删除失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("删除成功", c)
}

// GetNavBars 获取导航菜单列表
func GetNavBars(c *gin.Context) {
	navBars, err := services.GetNavBars()
	if err != nil {
		models.FailWithDetailed("", err.Error(), c)
		return
	}
	models.OkWithData(navBars, c)
}

// CreateNavBar 创建导航菜单
func CreateNavBar(c *gin.Context) {
	var navBar models.ProfileNavBar
	if err := c.ShouldBindJSON(&navBar); err != nil {
		models.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	if err := services.CreateNavBar(&navBar); err != nil {
		models.FailWithMessage("创建失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("创建成功", c)
}

// UpdateNavBar 更新导航菜单
func UpdateNavBar(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}
	var navBar models.ProfileNavBar
	if err := c.ShouldBindJSON(&navBar); err != nil {
		models.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	navBar.ID = uint(id)
	if err := services.UpdateNavBar(&navBar); err != nil {
		models.FailWithMessage("更新失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("更新成功", c)
}

// DeleteNavBar 删除导航菜单
func DeleteNavBar(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}
	if err := services.DeleteNavBar(uint(id)); err != nil {
		models.FailWithMessage("删除失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("删除成功", c)
}

// GetSkills 获取技能列表
func GetSkills(c *gin.Context) {
	data, err := services.GetSkills()
	if err != nil {
		models.FailWithDetailed("", err.Error(), c)
		return
	}
	models.OkWithData(data, c)
}

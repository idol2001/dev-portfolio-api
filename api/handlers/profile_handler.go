package handlers

import (
	"dev-portfolio-api/internal/services"
	"dev-portfolio-api/models"

	"github.com/gin-gonic/gin"
)

func SkillsHandler(c *gin.Context) {
	data, err := services.GetSkills()
	if err != nil {
		models.FailWithDetailed("", err.Error(), c)
		return
	}
	c.JSON(200, data)
}

func HomeHandler(c *gin.Context) {
	data, err := services.GetHome()
	if err != nil {
		models.FailWithDetailed("", err.Error(), c)
		return
	}
	c.JSON(200, data)
}

func AboutHandler(c *gin.Context) {
	data, err := services.GetAbout()
	if err != nil {
		models.FailWithDetailed("", err.Error(), c)
		return
	}
	c.JSON(200, data)
}

func NavBarHandler(c *gin.Context) {
	data, err := services.GetNavBar()
	if err != nil {
		models.FailWithDetailed("", err.Error(), c)
		return
	}
	c.JSON(200, data)
}

func SocialHandler(c *gin.Context) {
	data, err := services.GetSocial()
	if err != nil {
		models.FailWithDetailed("", err.Error(), c)
		return
	}
	c.JSON(200, data)
}

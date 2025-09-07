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

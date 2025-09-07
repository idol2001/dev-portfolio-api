package routes

import (
	"dev-portfolio-api/api/handlers"

	"github.com/gin-gonic/gin"
)

func InitProfileRouter(r *gin.RouterGroup) (R gin.IRoutes) {
	router := r.Group("profile")

	router.GET("/skills", handlers.SkillsHandler)

	return router
}

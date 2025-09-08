package routes

import (
	"dev-portfolio-api/api/handlers"

	"github.com/gin-gonic/gin"
)

func InitProfileRouter(r *gin.RouterGroup) (R gin.IRoutes) {
	router := r.Group("profile")

	router.GET("/skills", handlers.SkillsHandler)
	router.GET("/home", handlers.HomeHandler)
	router.GET("/about", handlers.AboutHandler)
	router.GET("/navbar", handlers.NavBarHandler)
	router.GET("/social", handlers.SocialHandler)

	return router
}

package routes

import (
	"dev-portfolio-api/api/handlers"
	"dev-portfolio-api/middleware"

	"github.com/gin-gonic/gin"
)

func InitProfileRouter(r *gin.RouterGroup) {
	// 认证路由（不需要登录）
	auth := r.Group("auth")
	{
		auth.POST("/login", handlers.Login)
		auth.GET("/public-key", handlers.GetPublicKey)
	}

	// 公开访问的路由（前台展示用）
	public := r.Group("")
	{
		public.GET("/profile/info", handlers.GetProfile)
		public.GET("/profile/skills", handlers.GetSkills)
		public.GET("/profile/socials", handlers.GetSocials)
		public.GET("/profile/navbar", handlers.GetNavBars)
		public.GET("/projects", handlers.GetProjects)
		public.GET("/projects/:id", handlers.GetProjectByID)
		public.GET("/blogs", handlers.GetBlogPosts)
		public.GET("/blogs/slug/:slug", handlers.GetBlogPostBySlug)
		// 静态文件服务（开发环境直接提供 uploads 目录）
		public.Static("/uploads", "./uploads")
	}

	// 需要认证的路由
	protected := r.Group("")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/user/me", handlers.GetCurrentUser)

		// 文件上传
		protected.POST("/upload/:type", handlers.UploadFile)

		profile := protected.Group("profile")
		{
			profile.PUT("/info", handlers.UpdateProfile)

			socials := profile.Group("socials")
			{
				socials.POST("", handlers.CreateSocial)
				socials.PUT("/:id", handlers.UpdateSocial)
				socials.DELETE("/:id", handlers.DeleteSocial)
			}

			navbar := profile.Group("navbar")
			{
				navbar.POST("", handlers.CreateNavBar)
				navbar.PUT("/:id", handlers.UpdateNavBar)
				navbar.DELETE("/:id", handlers.DeleteNavBar)
			}
		}

		projects := protected.Group("projects")
		{
			projects.POST("", handlers.CreateProject)
			projects.PUT("/:id", handlers.UpdateProject)
			projects.DELETE("/:id", handlers.DeleteProject)
		}

		blogs := protected.Group("blogs")
		{
			blogs.GET("/:id", handlers.GetBlogPostByID)
			blogs.POST("", handlers.CreateBlogPost)
			blogs.PUT("/:id", handlers.UpdateBlogPost)
			blogs.DELETE("/:id", handlers.DeleteBlogPost)
		}

		users := protected.Group("users")
		{
			users.GET("", handlers.GetUsers)
			users.POST("", handlers.CreateUser)
			users.PUT("/:id", handlers.UpdateUser)
			users.PUT("/:id/password", handlers.ChangePassword)
			users.DELETE("/:id", handlers.DeleteUser)
		}

		skills := protected.Group("skills")
		{
			skills.GET("", handlers.GetSkillGroups)
			skills.POST("/groups", handlers.CreateSkillGroup)
			skills.PUT("/groups/:id", handlers.UpdateSkillGroup)
			skills.DELETE("/groups/:id", handlers.DeleteSkillGroup)
			skills.GET("/groups/:groupId/items", handlers.GetSkillItems)
			skills.POST("/groups/:groupId/items", handlers.CreateSkillItem)
			skills.PUT("/items/:id", handlers.UpdateSkillItem)
			skills.DELETE("/items/:id", handlers.DeleteSkillItem)
		}
	}
}

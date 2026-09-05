package handlers

import (
	"dev-portfolio-api/internal/services"
	"dev-portfolio-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBlogPosts 获取博客列表
func GetBlogPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	published := c.DefaultQuery("published", "true")

	posts, total, err := services.GetBlogPosts(published == "true", page, pageSize)
	if err != nil {
		models.FailWithDetailed("", err.Error(), c)
		return
	}
	models.OkWithDataList(gin.H{
		"list":     posts,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}, int64(total), c)
}

// GetBlogPostBySlug 根据 slug 获取文章
func GetBlogPostBySlug(c *gin.Context) {
	slug := c.Param("slug")
	post, err := services.GetBlogPostBySlug(slug)
	if err != nil {
		models.FailWithDetailed("", err.Error(), c)
		return
	}
	models.OkWithData(post, c)
}

// GetBlogPostByID 根据 ID 获取文章
func GetBlogPostByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}
	post, err := services.GetBlogPostByID(uint(id))
	if err != nil {
		models.FailWithDetailed("", err.Error(), c)
		return
	}
	models.OkWithData(post, c)
}

// CreateBlogPost 创建文章
func CreateBlogPost(c *gin.Context) {
	var post models.BlogPost
	if err := c.ShouldBindJSON(&post); err != nil {
		models.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	// 自动填充当前用户ID为作者
	userId, exists := c.Get("userId")
	if exists {
		post.AuthorID = uint64(userId.(uint))
	}

	if err := services.CreateBlogPost(&post); err != nil {
		models.FailWithMessage("创建失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("创建成功", c)
}

// UpdateBlogPost 更新文章
func UpdateBlogPost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}
	var post models.BlogPost
	if err := c.ShouldBindJSON(&post); err != nil {
		models.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	post.ID = uint(id)
	if err := services.UpdateBlogPost(&post); err != nil {
		models.FailWithMessage("更新失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("更新成功", c)
}

// DeleteBlogPost 删除文章
func DeleteBlogPost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}
	if err := services.DeleteBlogPost(uint(id)); err != nil {
		models.FailWithMessage("删除失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("删除成功", c)
}

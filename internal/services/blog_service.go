package services

import (
	"dev-portfolio-api/models"
	"dev-portfolio-api/pkg/global"
	"time"
)

// GetBlogPosts 获取博客文章列表
func GetBlogPosts(publishedOnly bool, page, pageSize int) ([]models.BlogPost, int64, error) {
	posts := make([]models.BlogPost, 0)
	query := global.DB
	if publishedOnly {
		query = query.Where("status = ?", "published")
	}

	var total int64
	query.Model(&models.BlogPost{}).Count(&total)

	offset := (page - 1) * pageSize
	err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&posts).Error
	return posts, total, err
}

// GetBlogPostBySlug 根据 slug 获取文章
func GetBlogPostBySlug(slug string) (*models.BlogPost, error) {
	post := &models.BlogPost{}
	err := global.DB.Where("slug = ?", slug).First(post).Error
	if err != nil {
		return nil, err
	}
	// 增加阅读数
	global.DB.Model(post).UpdateColumn("view_count", post.ViewCount+1)
	return post, nil
}

// GetBlogPostByID 根据 ID 获取文章
func GetBlogPostByID(id uint) (*models.BlogPost, error) {
	post := &models.BlogPost{}
	err := global.DB.First(post, id).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

// CreateBlogPost 创建文章
func CreateBlogPost(post *models.BlogPost) error {
	if post.Status == "published" && post.PublishedAt == nil {
		now := time.Now().Unix()
		post.PublishedAt = &now
	}
	return global.DB.Create(post).Error
}

// UpdateBlogPost 更新文章
func UpdateBlogPost(post *models.BlogPost) error {
	if post.Status == "published" && post.PublishedAt == nil {
		now := time.Now().Unix()
		post.PublishedAt = &now
	}
	// 使用 Updates 只更新指定字段，避免 created_at 零值被写入
	return global.DB.Model(&models.BlogPost{}).Where("id = ?", post.ID).Updates(map[string]interface{}{
		"title":        post.Title,
		"slug":         post.Slug,
		"summary":      post.Summary,
		"content":      post.Content,
		"cover_image":  post.CoverImage,
		"status":       post.Status,
		"tags":         post.Tags,
		"published_at": post.PublishedAt,
	}).Error
}

// DeleteBlogPost 删除文章
func DeleteBlogPost(id uint) error {
	return global.DB.Delete(&models.BlogPost{}, id).Error
}

package models

// BlogPost 博客文章表
type BlogPost struct {
	BaseModel
	Title       string `gorm:"column:title;size:256;not null;comment:'文章标题'" json:"title"`
	Slug        string `gorm:"column:slug;size:256;uniqueIndex:idx_blog_posts_slug;not null;comment:'URL标识'" json:"slug"`
	Summary     string `gorm:"column:summary;size:512;comment:'文章摘要'" json:"summary"`
	Content     string `gorm:"column:content;type:longtext;comment:'文章内容(Markdown)'" json:"content"`
	CoverImage  string `gorm:"column:cover_image;size:512;comment:'封面图URL'" json:"cover_image"`
	AuthorID    uint64 `gorm:"column:author_id;not null;comment:'作者ID'" json:"author_id"`
	Status      string `gorm:"column:status;size:32;default:'draft';comment:'状态(draft/published)'" json:"status"`
	Tags        string `gorm:"column:tags;size:512;comment:'标签(逗号分隔)'" json:"tags"`
	PublishedAt *int64 `gorm:"column:published_at;comment:'发布时间(Unix时间戳)'" json:"published_at"`
	ViewCount   int64  `gorm:"column:view_count;default:0;comment:'浏览次数'" json:"view_count"`
}

func (BlogPost) TableName() string {
	return "blog_posts"
}

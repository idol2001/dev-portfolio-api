package models

// Project 项目表
type Project struct {
	BaseModel
	Title       string `gorm:"column:title;size:256;not null;comment:'项目标题'" json:"title"`
	Description string `gorm:"column:description;type:longtext;comment:'项目描述'" json:"description"`
	TechStack   string `gorm:"column:tech_stack;type:text;comment:'技术栈(JSON数组)'" json:"tech_stack"`
	RepoURL     string `gorm:"column:repo_url;size:512;comment:'仓库URL'" json:"repo_url"`
	DemoURL     string `gorm:"column:demo_url;size:512;comment:'在线Demo URL'" json:"demo_url"`
	CoverImage  string `gorm:"column:cover_image;size:512;comment:'封面图URL'" json:"cover_image"`
	OrderNo     int64  `gorm:"column:order_no;default:0;index:idx_projects_order_no;comment:'排序号'" json:"order_no"`
	Status      string `gorm:"column:status;size:32;default:'published';comment:'状态(draft/published)'" json:"status"`
}

func (Project) TableName() string {
	return "profile_projects"
}

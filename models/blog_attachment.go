package models

// BlogAttachment 博客附件表（OSS 关联）
type BlogAttachment struct {
	BaseModel
	PostID   uint64 `gorm:"column:post_id;index" json:"post_id"`
	OSSKey   string `gorm:"column:oss_key;size:512;not null" json:"oss_key"`
	OSSURL   string `gorm:"column:oss_url;size:512;not null" json:"oss_url"`
	FileName string `gorm:"column:file_name;size:256" json:"file_name"`
	FileSize int64  `gorm:"column:file_size;default:0" json:"file_size"`
	FileType string `gorm:"column:file_type;size:64" json:"file_type"`
}

func (BlogAttachment) TableName() string {
	return "blog_attachments"
}

package model

type Comment struct {
	Base
	ArticleId uint   `json:"article_id;comment:文章主键"`
	Username  string `json:"size:256;comment:用户名"`
	Email     string `json:"size:256;comment:电子邮箱"`
	Content   string `gorm:"type:varchar(200);not null;comment:内容" json:"content"`
	Status    int8   `gorm:"type:tinyint;default:2;comment:状态" json:"status"`
}

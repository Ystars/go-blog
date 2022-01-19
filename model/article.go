package model

// Article 文章
type Article struct {
	Init
	State        int    `gorm:"type:tinyint;not null;default:1;comment:状态" json:"state" label:"状态"`
	IsTop        int    `gorm:"type:tinyint;not null;default:0;comment:是否置顶" json:"is_top" label:"是否置顶"`
	Title        string `gorm:"type:varchar(100);not null;comment:标题" json:"title" label:"标题"`
	Desc         string `gorm:"type:varchar(200);comment:描述" json:"desc" label:"描述"`
	Content      string `gorm:"type:longtext;comment:内容" json:"content" label:"内容"`
	CommentCount int    `gorm:"type:int;not null;default:0;comment:评论次数" json:"comment_count" label:"评论次数"`
	ReadCount    int    `gorm:"type:int;not null;default:0;comment:阅读次数" json:"read_count" label:"阅读次数"`
}

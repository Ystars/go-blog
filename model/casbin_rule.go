package model

// CasbinRule 文章
type CasbinRule struct {
	Base
	Ptype     string `gorm:"size:100;uniqueIndex:unique_index" json:"ptype"`
	Authority string `gorm:"size:100;uniqueIndex:unique_index;column:v0" json:"rolename"`
	Path      string `gorm:"size:100;uniqueIndex:unique_index;column:v1" json:"v1"`
	Method    string `gorm:"size:100;uniqueIndex:unique_index;column:v2" json:"v2"`
}

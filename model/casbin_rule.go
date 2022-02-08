package model

// CasbinRule 文章
type CasbinRule struct {
	Base
	Ptype     string `gorm:"size:100;uniqueIndex:unique_index"`
	Authority string `gorm:"size:100;uniqueIndex:unique_index;column:v0"`
	Path      string `gorm:"size:100;uniqueIndex:unique_index;column:v1"`
	Method    string `gorm:"size:100;uniqueIndex:unique_index;column:v2"`
	V3        string `gorm:"size:100;uniqueIndex:unique_index;"`
	V4        string `gorm:"size:100;uniqueIndex:unique_index;"`
	V5        string `gorm:"size:100;uniqueIndex:unique_index;"`
}

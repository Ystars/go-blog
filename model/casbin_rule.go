package model

// CasbinRule 文章
type CasbinRule struct {
	Base
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Ptype string `gorm:"size:100;uniqueIndex:unique_index" json:"p_type"`
	V0    string `gorm:"size:100;uniqueIndex:unique_index" json:"v0"`
	V1    string `gorm:"size:100;uniqueIndex:unique_index" json:"v1"`
	V2    string `gorm:"size:100;uniqueIndex:unique_index" json:"v2"`
	V3    string `gorm:"size:100;uniqueIndex:unique_index" json:"v3"`
	V4    string `gorm:"size:100;uniqueIndex:unique_index" json:"v4"`
	V5    string `gorm:"size:100;uniqueIndex:unique_index" json:"v5"`
}

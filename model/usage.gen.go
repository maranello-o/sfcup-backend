// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUsage = "usage"

// Usage mapped from table <usage>
type Usage struct {
	ID         int64  `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	UserID     int64  `gorm:"column:user_id;type:bigint;not null" json:"user_id"`
	Model      string `gorm:"column:model;type:varchar(64);not null" json:"model"`
	CreateTime int64  `gorm:"autoCreateTime" json:"create_time"`
}

// TableName Usage's table name
func (*Usage) TableName() string {
	return TableNameUsage
}

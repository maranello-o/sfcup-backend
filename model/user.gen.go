// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID         int64  `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	Email      string `gorm:"column:email;type:varchar(64);not null" json:"email"`
	Password   string `gorm:"column:password;type:varchar(64);not null" json:"password"`
	Nickname   string `gorm:"column:nickname;type:varchar(32);not null" json:"nickname"`
	CreateTime int64  `gorm:"autoCreateTime" json:"create_time"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `grom:"uniqueIndex;type:varchar(128) not null"`	//存储登录屏障（唯一）
	Password string `grom:"type:varchar(64) not null"`	//存储密码
}

func (User) TableName() string {
	return "user"
}
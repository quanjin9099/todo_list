package main

import "time"

type TodoInfo struct {
	Id          int    `gorm:"primary_key;auto_increment"` //ID 主键 自增
	Name        string `gorm:"type:varchar(32)"`           //名称 最多32个字符
	Description string `gorm:"type:varchar(255)"`          //描述 最多255个字符
	Score       int    `gorm:"type:int"`                   //分数 1-100
	Status      int    `gorm:"type:int"`                   //状态 0 未完成 1 已完成
	UserID      int    `gorm:"type:int"`                   //用户ID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type User struct {
	Id   int    `gorm:"primary_key;auto_increment"` //ID 主键 自增
	Name string `gorm:"type:varchar(32)"`           //名称 最多32个字符
}

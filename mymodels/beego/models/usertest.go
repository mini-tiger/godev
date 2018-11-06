package models

import "time"

type User struct {
	Id int
	Name string
	Update time.Time
}

//自定义表名
func (u *User) TableName() string {
	return "auth_user"
}

// 多字段索引
func (u *User) TableIndex() [][]string {
	return [][]string{
		[]string{"Id", "Name"},
	}
}

// 多字段唯一键
func (u *User) TableUnique() [][]string {
	return [][]string{
		[]string{"Name", "Email"},
	}
}

// 设置引擎为 INNODB
func (u *User) TableEngine() string {
	return "INNODB"
}
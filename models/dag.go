package models

import (
	"time"
)

type Dag struct {
	Id         int       `xorm:"not null pk autoincr comment('dag ID') INT"`
	Name       string    `xorm:"comment('名称') VARCHAR(128)"`
	Config     string    `xorm:"comment('dag 配置') JSON"`
	Status     int       `xorm:"comment('dag 状态') INT"`
	CreateTime time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateTime time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
}

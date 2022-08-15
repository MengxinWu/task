package models

import (
	"time"
)

type Processor struct {
	Id         int       `xorm:"not null pk autoincr comment('processor ID') INT"`
	Name       string    `xorm:"comment('名称') VARCHAR(128)"`
	Handler    string    `xorm:"comment('处理器(英文)') VARCHAR(128)"`
	Status     int       `xorm:"comment('处理器 状态') INT"`
	CreateTime time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateTime time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
}

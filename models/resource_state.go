package models

import (
	"time"
)

type ResourceState struct {
	Id           int       `xorm:"not null pk autoincr comment('id') INT"`
	ResourceId   int64     `xorm:"comment('资源 id') BIGINT"`
	ProcessorId  int       `xorm:"comment('处理器 id') INT"`
	ProcessState int       `xorm:"comment('处理状态') INT"`
	CreateTime   time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateTime   time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
}

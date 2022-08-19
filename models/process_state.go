package models

import (
	"time"
)

type ResourceProcessState struct {
	Id           int       `xorm:"not null pk autoincr comment('id') INT"`
	ResourceId   int64     `xorm:"comment('资源 id') BIGINT"`
	ProcessorId  int       `xorm:"comment('处理器 id') INT"`
	ProcessCnt   int       `xorm:"comment('处理次数') INT"`
	ProcessState int       `xorm:"comment('处理状态') INT"`
	ProcessMsg   string    `xorm:"comment('处理状态') VARCHAR(512)"`
	CreateTime   time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateTime   time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
}

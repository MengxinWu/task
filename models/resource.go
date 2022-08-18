package models

import (
	"time"
)

type Resource struct {
	ResourceId int64     `xorm:"not null pk comment('资源 id') BIGINT"`
	DagId      int       `xorm:"not null comment('Graph id') INT"`
	Name       string    `xorm:"comment('名称') VARCHAR(128)"`
	Status     int       `xorm:"comment('状态') INT"`
	CreateTime time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateTime time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
}

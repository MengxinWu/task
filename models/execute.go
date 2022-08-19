package models

type ExecuteEvent struct {
	ResourceId  int64 `json:"resource_id"`
	ProcessorId int   `json:"processor_id"`
	// 处理中间变量
	Resource      *Resource      `json:"-"`
	Processor     *Processor     `json:"-"`
	ResourceState *ResourceState `json:"-"`
	// 处理状态
	ProcessState int `json:"-"`
}

package models

type ExecuteEvent struct {
	ResourceId  int64 `json:"resource_id"`
	ProcessorId int   `json:"processor_id"`
	// 处理中间变量
	Resource             *Resource             `json:"-"`
	Processor            *Processor            `json:"-"`
	ResourceProcessState *ResourceProcessState `json:"-"`
	// 处理结果
	ProcessMsg   string `json:"-"`
	ProcessCnt   int    `json:"-"`
	ProcessState int    `json:"-"`
}

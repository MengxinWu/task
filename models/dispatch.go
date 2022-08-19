package models

const (
	DispatchEventResourceAdd   = "resource_add"
	DispatchEventProcessorDone = "processor_done" // 任务完成包含任务成功和失败
)

type DispatchEvent struct {
	// 初始化变量
	Event       string `json:"event"`
	ResourceId  int64  `json:"resource_id"`
	DagId       int    `json:"dag_id"`
	ProcessorId int    `json:"processor_id"`
	// 中间计算使用变量
	Resource             *Resource             `json:"-"`
	Dag                  *Dag                  `json:"-"`
	ResourceProcessState *ResourceProcessState `json:"-"`
	Graph                Graph                 `json:"-"`
	// 调度结果
	ExecutorList []int64 `json:"-"`
}

package models

const (
	EventResourceAdd = "resource_add"
)

type DispatchEvent struct {
	//
	Event       string `json:"event"`
	ResourceId  int64  `json:"resource_id"`
	DagId       int    `json:"dag_id"`
	ProcessorId int    `json:"processor_id"`
	//
	Resource *Resource `json:"-"`
	Dag      *Dag      `json:"dag"`
	Graph    Graph     `json:"graph"`

	//
	ExecutorList []int64 `json:"-"`
}

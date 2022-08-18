package models

const (
	DispatchEventResourceAdd = "resource_add"
)

type DispatchEvent struct {
	//
	Event       string `json:"event"`
	ResourceId  int64  `json:"resource_id"`
	DagId       int    `json:"dag_id"`
	ProcessorId int    `json:"processor_id"`
	//
	Resource *Resource `json:"-"`
	Dag      *Dag      `json:"-"`
	Graph    Graph     `json:"-"`

	//
	ExecutorList []int64 `json:"-"`
}

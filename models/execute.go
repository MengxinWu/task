package models

type ExecuteEvent struct {
	ResourceId   int64 `json:"resource_id"`
	ProcessorId  int   `json:"processor_id"`
	ProcessState int   `json:"-"`
}

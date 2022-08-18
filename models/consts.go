package models

const (
	DagStatusDelete = -1

	ResourceStatusNormal = 0
	ResourceStatusDELETE = -1

	ProcessStateInit    = 0
	ProcessStateReady   = 100
	ProcessStateRunning = 200
	ProcessStateFinish  = 400
	ProcessStateFail    = 500
	ProcessStateError   = 600
)

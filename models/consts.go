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

const (
	KafkaAddress          = "kafka:9092"
	KafkaTopicDispatch    = "dispatch-topic"
	KafkaTopicExecute     = "execute-topic"
	KafkaConsumerDispatch = "consumer-group-dispatch"
	KafkaConsumerExecute  = "consumer-group-execute"
)

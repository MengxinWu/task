package models

const (
	// resource
	ResourceStatusNormal = 0
	ResourceStatusDelete = -1

	// dag
	DagStatusDelete = -1

	// processor
	ProcessorStatusNormal = int(0)
	ProcessorStatusDelete = int(-1)

	ProcessStateInit    = 0
	ProcessStateReady   = 100
	ProcessStateRunning = 200
	ProcessStateSuccess = 400
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

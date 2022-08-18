package models

const (
	DagStatusDelete = -1

	// resource
	ResourceStatusNormal = 0
)

const (
	KafkaAddress          = "kafka:9092"
	KafkaTopicDispatch    = "dispatch-topic"
	KafkaTopicExecute     = "execute-topic"
	KafkaConsumerDispatch = "consumer-group-dispatch"
	KafkaConsumerExecute  = "consumer-group-execute"
)

package model

type Event interface {
	GetTopic() Topic
}

type Topic string

const (
	SKIP_TOPIC = "skip_topic"
)

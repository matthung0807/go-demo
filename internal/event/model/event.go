package model

type Event interface {
	GetTopic() Topic
}

type Topic string

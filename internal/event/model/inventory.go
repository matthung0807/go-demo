package model

import "abc.com/demo/internal/domain"

const (
	CREATE_INVENTORY_TOPIC       Topic = "create_inventory_topic"
	CREATE_INVENTORY_REPLY_TOPIC Topic = "create_inventory_reply_topic"

	UPDATE_INVENTORY_TOPIC       Topic = "update_inventory_topic"
	UPDATE_INVENTORY_REPLY_TOPIC Topic = "update_inventory_reply_topic"
)

type CreateInventoryEvent struct {
	EventMessage[CreateInventoryPayload]
}

func NewCreateInventoryEvent(corId string, payload CreateInventoryPayload) CreateInventoryEvent {
	return CreateInventoryEvent{
		EventMessage: NewEventMessage(corId, CREATE_INVENTORY_TOPIC, payload),
	}
}

type CreateInventoryPayload struct {
	UserId string
}

type CreateInventoryReplyEvent struct {
	EventMessage[CreateInventoryReplyPayload]
}

func NewCreateInventoryReplyEvent(corId string, payload CreateInventoryReplyPayload) CreateInventoryReplyEvent {
	return CreateInventoryReplyEvent{
		EventMessage: NewEventMessage(corId, CREATE_INVENTORY_REPLY_TOPIC, payload),
	}
}

type CreateInventoryReplyPayload struct {
	Inventory domain.Inventory
	Result    ReplyResult
	UserId    string
}

type UpdateInventoryEvent struct {
	EventMessage[UpdateInventoryPayload]
}

func NewUpdateInventoryEvent(corId string, payload UpdateInventoryPayload) UpdateInventoryEvent {
	return UpdateInventoryEvent{
		EventMessage: NewEventMessage(corId, UPDATE_INVENTORY_TOPIC, payload),
	}
}

type UpdateInventoryPayload struct {
	UserId string
}

type UpdateInventoryReplyEvent struct {
	EventMessage[UpdateInventoryReplyPayload]
}

func NewUpdateInventoryReplyEvent(corId string, payload UpdateInventoryReplyPayload) UpdateInventoryReplyEvent {
	return UpdateInventoryReplyEvent{
		EventMessage: NewEventMessage(corId, UPDATE_INVENTORY_REPLY_TOPIC, payload),
	}
}

type UpdateInventoryReplyPayload struct {
	Inventory domain.Inventory
	Result    ReplyResult
	UserId    string
}

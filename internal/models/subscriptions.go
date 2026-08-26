package models

// Subscription models the NimbleOS subscriptions object set.
type Subscription struct {
	// ID Identifier for subscription.
	ID string `json:"id,omitempty"`
	// SubscriberID Identifier for subscriber (notification client) that this subscription belongs to.
	SubscriberID string `json:"subscriber_id,omitempty"`
	// NotificationType This indicates the type of notification being subscribed for.
	NotificationType string `json:"notification_type,omitempty"`
	// ObjectType The object type that the notification subscriber is interested in.
	ObjectType string `json:"object_type,omitempty"`
	// ObjectID The object that the notification subscriber is interested in.
	ObjectID string `json:"object_id,omitempty"`
	// Operation The operation that the notification subscriber is interested in.
	Operation float64 `json:"operation,omitempty"`
	// EventTargetType The kind of events or alerts that the notification subscriber is interested in.
	EventTargetType string `json:"event_target_type,omitempty"`
	// EventSeverity The severity of events that the notification subscriber is interested in.
	EventSeverity string `json:"event_severity,omitempty"`
}

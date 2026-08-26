package models

// Alarm models the NimbleOS alarms object set.
type Alarm struct {
	// ID Identifier for the alarm.
	ID string `json:"id,omitempty"`
	// Type Identifier for type of alarm.
	Type string `json:"type,omitempty"`
	// Array The array name where the alarm is generated.
	Array string `json:"array,omitempty"`
	// CurrOnsetEventID Identifier for the current onset event.
	CurrOnsetEventID string `json:"curr_onset_event_id,omitempty"`
	// ObjectID Identifier of object operated upon.
	ObjectID string `json:"object_id,omitempty"`
	// ObjectName Name of object operated upon.
	ObjectName string `json:"object_name,omitempty"`
	// ObjectType Type of the object being operated upon.
	ObjectType string `json:"object_type,omitempty"`
	// OnsetTime Time when this alarm was triggered.
	OnsetTime int64 `json:"onset_time,omitempty"`
	// AckTime Time when this alarm was acknowledged.
	AckTime int64 `json:"ack_time,omitempty"`
	// Status Status of the operation -- open or acknowledged.
	Status string `json:"status,omitempty"`
	// UserID Identifier of the user who acknowledged the alarm.
	UserID string `json:"user_id,omitempty"`
	// UserName Username of the user who acknowledged the alarm.
	UserName string `json:"user_name,omitempty"`
	// UserFullName Full name of the user who acknowledged the alarm.
	UserFullName string `json:"user_full_name,omitempty"`
	// Category Category of the alarm.
	Category string `json:"category,omitempty"`
	// Severity Severity level of the event.
	Severity string `json:"severity,omitempty"`
	// RemindEvery Frequency of notification.
	RemindEvery string `json:"remind_every,omitempty"`
	// RemindEveryUnit Time unit over which to send the number of notification specified in 'remind_every'.
	RemindEveryUnit string `json:"remind_every_unit,omitempty"`
	// Activity Description of activity performed and recorded in alarm.
	Activity string `json:"activity,omitempty"`
	// NextNotificationTime Time when next reminder for the alarm will be sent.
	NextNotificationTime int64 `json:"next_notification_time,omitempty"`
}

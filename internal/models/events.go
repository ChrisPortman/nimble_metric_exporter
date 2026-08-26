package models

// Event models the NimbleOS events object set.
type Event struct {
	// ID Identifier for the event record.
	ID string `json:"id,omitempty"`
	// Type Type of the event record.
	Type string `json:"type,omitempty"`
	// Name Name of alert macro to generate.
	Name string `json:"name,omitempty"`
	// Scope The array name for array level event.
	Scope string `json:"scope,omitempty"`
	// Target Name of object upon which the event occurred.
	Target string `json:"target,omitempty"`
	// TargetType Target type of the event record.
	TargetType string `json:"target_type,omitempty"`
	// Timestamp Time when this event happened.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Category Category of the event record.
	Category string `json:"category,omitempty"`
	// Severity Severity level of the event.
	Severity string `json:"severity,omitempty"`
	// Summary Summary of the event.
	Summary string `json:"summary,omitempty"`
	// Activity Description of the event.
	Activity string `json:"activity,omitempty"`
	// AlarmID The alarm ID if the event is related to an alarm.
	AlarmID string `json:"alarm_id,omitempty"`
	// Params Arguments provided for event creation in key-value structure.
	Params []any `json:"params,omitempty"`
	// TenantID Tenant ID of the event.
	TenantID string `json:"tenant_id,omitempty"`
}

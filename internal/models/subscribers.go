package models

// Subscriber models the NimbleOS subscribers object set.
type Subscriber struct {
	// ID Identifier for subscriber.
	ID string `json:"id,omitempty"`
	// Type This is generally used to indicate the type of subscriber e.
	Type string `json:"type,omitempty"`
	// RenewInterval The interval in seconds within which the subscriber is expected to send a renew message over the websocket channel in case there is no traffic on the websocket channel.
	RenewInterval int64 `json:"renew_interval,omitempty"`
	// RenewResponseTimeout The interval in seconds after the subscriber sends a renew message within which the subscriber expects to get a response.
	RenewResponseTimeout int64 `json:"renew_response_timeout,omitempty"`
	// IsConnected True if the subscriber has an active websocket connection.
	IsConnected bool `json:"is_connected,omitempty"`
	// NotificationCount Number of notifications sent to subscriber.
	NotificationCount int64 `json:"notification_count,omitempty"`
	// Force Forcibly modify a connected subscriber.
	Force bool `json:"force,omitempty"`
}

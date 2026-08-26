package models

// KeyManager models the NimbleOS key_managers object set.
type KeyManager struct {
	// ID Identifier for External Key Manager.
	ID string `json:"id,omitempty"`
	// Name Name of external key manager.
	Name string `json:"name,omitempty"`
	// Description Description of external key manager.
	Description string `json:"description,omitempty"`
	// Hostname Hostname or IP Address for the External Key Manager.
	Hostname string `json:"hostname,omitempty"`
	// Port Port number for the External Key Manager.
	Port int64 `json:"port,omitempty"`
	// Protocol KMIP protocol supported by External Key Manager.
	Protocol string `json:"protocol,omitempty"`
	// Username External Key Manager username.
	Username string `json:"username,omitempty"`
	// Password External Key Manager user password.
	Password string `json:"password,omitempty"`
	// Active Whether the given key manager is active or not.
	Active bool `json:"active,omitempty"`
	// Status Connection status of a given external key manager.
	Status string `json:"status,omitempty"`
	// Vendor KMIP vendor name.
	Vendor string `json:"vendor,omitempty"`
}

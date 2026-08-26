package models

// Witness models the NimbleOS witnesses object set.
type Witness struct {
	// ID Identifier of the witness configuration.
	ID string `json:"id,omitempty"`
	// Username Username of witness.
	Username string `json:"username,omitempty"`
	// Password Password of witness.
	Password string `json:"password,omitempty"`
	// Host Hostname or ip addresses of witness.
	Host string `json:"host,omitempty"`
	// Port Port of witness.
	Port int64 `json:"port,omitempty"`
	// SecureMode To verify the witness host against CA cert and to apply possible security modes.
	SecureMode bool `json:"secure_mode,omitempty"`
	// AutoSwitchoverMessages List of validation messages for automatic switchover of Group Management.
	AutoSwitchoverMessages []any `json:"auto_switchover_messages,omitempty"`
}

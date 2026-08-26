package models

// MasterKey models the NimbleOS master_key object set.
type MasterKey struct {
	// ID Identifier of the master key.
	ID string `json:"id,omitempty"`
	// Name Name of the master key.
	Name string `json:"name,omitempty"`
	// Passphrase Passphrase used to protect the master key, required during creation, enabling/disabling the key and change the passphrase to a new value.
	Passphrase string `json:"passphrase,omitempty"`
	// Halfkey When changing the passphrase, this authenticates the change operation, for support use only.
	Halfkey string `json:"halfkey,omitempty"`
	// NewPassphrase When changing the passphrase, this attribute specifies the new value of the passphrase.
	NewPassphrase string `json:"new_passphrase,omitempty"`
	// Active Whether the master key is active or not.
	Active bool `json:"active,omitempty"`
	// PurgeAge Default minimum age (in hours) of inactive encryption keys to be purged.
	PurgeAge int64 `json:"purge_age,omitempty"`
}

package models

// ClientCredential models the NimbleOS client_credentials object set.
type ClientCredential struct {
	// ID Identifier for the client credentials record.
	ID string `json:"id,omitempty"`
	// Name Client name.
	Name string `json:"name,omitempty"`
	// ClientID Uniqely identify this client, preferably uuid.
	ClientID string `json:"client_id,omitempty"`
	// Secret Client secret corresponding to this client id and name.
	Secret string `json:"secret,omitempty"`
	// CreationTime Time when this client credentials was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this client credentials was last modified.
	LastModified string `json:"last_modified,omitempty"`
	// Description Description of client.
	Description string `json:"description,omitempty"`
}

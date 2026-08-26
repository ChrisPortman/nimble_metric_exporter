package models

// ApplicationCategory models the NimbleOS application_categories object set.
type ApplicationCategory struct {
	// ID Identifier for the application category.
	ID string `json:"id,omitempty"`
	// Name Name of application category.
	Name string `json:"name,omitempty"`
	// DedupeEnabled Specifies if dedupe is enabled for performance policies associated with this application category.
	DedupeEnabled bool `json:"dedupe_enabled,omitempty"`
	// CreationTime Time when this application category was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this application category was last modified.
	LastModified string `json:"last_modified,omitempty"`
}

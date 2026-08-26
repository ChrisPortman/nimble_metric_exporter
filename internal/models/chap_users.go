package models

// ChapUser models the NimbleOS chap_users object set.
type ChapUser struct {
	// ID Identifier for the CHAP user.
	ID string `json:"id,omitempty"`
	// Name Name of CHAP user.
	Name string `json:"name,omitempty"`
	// FullName CHAP user's fully qualified name.
	FullName string `json:"full_name,omitempty"`
	// SearchName CHAP user name used for object search.
	SearchName string `json:"search_name,omitempty"`
	// Description Text description of CHAP user.
	Description string `json:"description,omitempty"`
	// Password CHAP secret.
	Password string `json:"password,omitempty"`
	// InitiatorIqns List of iSCSI initiators.
	InitiatorIqns []string `json:"initiator_iqns,omitempty"`
	// TenantID Identifier for the tenant.
	TenantID string `json:"tenant_id,omitempty"`
	// CreationTime Time when this CHAP user was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this CHAP user was last modified.
	LastModified string `json:"last_modified,omitempty"`
	// VolList List of volumes associated with this CHAP user.
	VolList []any `json:"vol_list,omitempty"`
	// VolCount Count of volumes associated with this CHAP user.
	VolCount int64 `json:"vol_count,omitempty"`
}

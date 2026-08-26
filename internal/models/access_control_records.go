package models

// AccessControlRecord models the NimbleOS access_control_records object set.
type AccessControlRecord struct {
	// ID Identifier for the access control record.
	ID string `json:"id,omitempty"`
	// ApplyTo Type of object this access control record applies to.
	ApplyTo string `json:"apply_to,omitempty"`
	// ChapUserID Identifier for the CHAP user.
	ChapUserID string `json:"chap_user_id,omitempty"`
	// ChapUserName Name of the CHAP user.
	ChapUserName string `json:"chap_user_name,omitempty"`
	// InitiatorGroupID Identifier for the initiator group.
	InitiatorGroupID string `json:"initiator_group_id,omitempty"`
	// InitiatorGroupName Name of the initiator group.
	InitiatorGroupName string `json:"initiator_group_name,omitempty"`
	// Lun If this access control record applies to a regular volume, this attribute is the volume's LUN (Logical Unit Number).
	Lun int64 `json:"lun,omitempty"`
	// VolID Identifier for the volume this access control record applies to.
	VolID string `json:"vol_id,omitempty"`
	// VolName Name of the volume this access control record applies to.
	VolName string `json:"vol_name,omitempty"`
	// VolAgentType External management agent type.
	VolAgentType int64 `json:"vol_agent_type,omitempty"`
	// PeID Identifier for the protocol endpoint this access control record applies to.
	PeID string `json:"pe_id,omitempty"`
	// PeName Name of the protocol endpoint this access control record applies to.
	PeName string `json:"pe_name,omitempty"`
	// PeLun LUN (Logical Unit Number) to associate with this protocol endpoint.
	PeLun int64 `json:"pe_lun,omitempty"`
	// SnapID Identifier for the snapshot this access control record applies to.
	SnapID string `json:"snap_id,omitempty"`
	// SnapName Name of the snapshot this access control record applies to.
	SnapName string `json:"snap_name,omitempty"`
	// PeIds List of candidate protocol endpoints that may be used to access the Virtual Volume.
	PeIds []any `json:"pe_ids,omitempty"`
	// Snapluns Information about the snapshot LUNs associated with this access control record.
	Snapluns []any `json:"snapluns,omitempty"`
	// CreationTime Time when this access control record was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this access control record was last modified.
	LastModified string `json:"last_modified,omitempty"`
	// AccessProtocol Access protocol of the volume.
	AccessProtocol string `json:"access_protocol,omitempty"`
}

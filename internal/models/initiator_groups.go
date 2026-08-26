package models

// InitiatorGroup models the NimbleOS initiator_groups object set.
type InitiatorGroup struct {
	// ID Identifier for initiator group.
	ID string `json:"id,omitempty"`
	// Name Name of initiator group.
	Name string `json:"name,omitempty"`
	// FullName Initiator group's full name.
	FullName string `json:"full_name,omitempty"`
	// SearchName Initiator group name used for search.
	SearchName string `json:"search_name,omitempty"`
	// Description Text description of initiator group.
	Description string `json:"description,omitempty"`
	// AccessProtocol Initiator group access protocol.
	AccessProtocol string `json:"access_protocol,omitempty"`
	// HostType Initiator group host type.
	HostType string `json:"host_type,omitempty"`
	// FCTdzPorts List of target Fibre Channel ports with Target Driven Zoning configured on this initiator group.
	FCTdzPorts []string `json:"fc_tdz_ports,omitempty"`
	// TargetSubnets List of target subnet labels.
	TargetSubnets []string `json:"target_subnets,omitempty"`
	// ISCSIInitiators List of iSCSI initiators.
	ISCSIInitiators []any `json:"iscsi_initiators,omitempty"`
	// FCInitiators List of FC initiators.
	FCInitiators []any `json:"fc_initiators,omitempty"`
	// CreationTime Time when this initiator group was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this initiator group was last modified.
	LastModified string `json:"last_modified,omitempty"`
	// VpOverride Flag to allow modifying VP created initiator groups.
	VpOverride bool `json:"vp_override,omitempty"`
	// AppUUID Application identifier of initiator group.
	AppUUID string `json:"app_uuid,omitempty"`
	// VolumeCount Number of volumes that are accessible by the initiator group.
	VolumeCount int64 `json:"volume_count,omitempty"`
	// VolumeList List of volumes that are accessible by the initiator group.
	VolumeList []any `json:"volume_list,omitempty"`
	// NumConnections Total number of connections from initiators in the initiator group.
	NumConnections string `json:"num_connections,omitempty"`
	// Metadata Key-value pairs that augment an initiator group's attributes.
	Metadata []any `json:"metadata,omitempty"`
}

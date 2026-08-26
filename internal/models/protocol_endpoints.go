package models

// ProtocolEndpoint models the NimbleOS protocol_endpoints object set.
type ProtocolEndpoint struct {
	// ID Identifier for the protocol endpoint.
	ID string `json:"id,omitempty"`
	// Name Name of the protocol endpoint.
	Name string `json:"name,omitempty"`
	// Description Text description of the protocol endpoint.
	Description string `json:"description,omitempty"`
	// PoolName Name of the pool where the protocol endpoint resides.
	PoolName string `json:"pool_name,omitempty"`
	// PoolID Identifier associated with the pool in the storage pool table.
	PoolID string `json:"pool_id,omitempty"`
	// State Operational state of protocol endpoint.
	State string `json:"state,omitempty"`
	// SerialNumber Identifier associated with the protocol endpoint for the SCSI protocol.
	SerialNumber int64 `json:"serial_number,omitempty"`
	// TargetName The iSCSI Qualified Name (IQN) or the Fibre Channel World Wide Node Name (WWNN) of the target protocol endpoint.
	TargetName string `json:"target_name,omitempty"`
	// GroupSpecificIds External UID is used to compute the serial number and IQN which never change even if the running group changes (e.
	GroupSpecificIds string `json:"group_specific_ids,omitempty"`
	// CreationTime Time when this protocol endpoint was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this protocol endpoint was last modified.
	LastModified string `json:"last_modified,omitempty"`
	// NumConnections Number of connections via this protocol endpoint.
	NumConnections string `json:"num_connections,omitempty"`
	// NumISCSIConnections Number of iSCSI connections via this protocol endpoint.
	NumISCSIConnections string `json:"num_iscsi_connections,omitempty"`
	// NumFCConnections Number of FC connections via this protocol endpoint.
	NumFCConnections string `json:"num_fc_connections,omitempty"`
	// AccessControlRecords List of access control records that apply to this protocol endpoint.
	AccessControlRecords []any `json:"access_control_records,omitempty"`
	// ISCSISessions List of iSCSI sessions connected to this protocol endpoint.
	ISCSISessions []any `json:"iscsi_sessions,omitempty"`
	// FCSessions List of FC sessions connected to this protocol endpoint.
	FCSessions []any `json:"fc_sessions,omitempty"`
	// AccessProtocol Access protocol of the protocol endpoint.
	AccessProtocol string `json:"access_protocol,omitempty"`
}

package models

// Snapshot models the NimbleOS snapshots object set.
type Snapshot struct {
	// ID Identifier for the snapshot.
	ID string `json:"id,omitempty"`
	// Name Name of snapshot.
	Name string `json:"name,omitempty"`
	// Description Text description of snapshot.
	Description string `json:"description,omitempty"`
	// Size Size of volume at time of snapshot (in bytes).
	Size int64 `json:"size,omitempty"`
	// VolName Name of the parent volume in which the snapshot belongs to.
	VolName string `json:"vol_name,omitempty"`
	// PoolName Name of the pool in which the parent volume belongs to.
	PoolName string `json:"pool_name,omitempty"`
	// VolID Parent volume ID.
	VolID string `json:"vol_id,omitempty"`
	// SnapCollectionName Name of snapshot collection.
	SnapCollectionName string `json:"snap_collection_name,omitempty"`
	// SnapCollectionID Identifier of snapshot collection.
	SnapCollectionID string `json:"snap_collection_id,omitempty"`
	// Online Online state for a snapshot means it could be mounted for data restore.
	Online bool `json:"online,omitempty"`
	// Writable Allow snapshot to be writable.
	Writable bool `json:"writable,omitempty"`
	// OfflineReason Snapshot offline reason - possible entries: one of 'user', 'recovery', 'replica', 'over_volume_limit', 'over_snapshot_limit', 'over_volume_reserve', 'nvram_loss_recovery', 'pool_free_space_exhausted'.
	OfflineReason string `json:"offline_reason,omitempty"`
	// ExpiryTime Unix timestamp indicating that the snapshot is considered expired by Snapshot Time-to-live(TTL).
	ExpiryTime int64 `json:"expiry_time,omitempty"`
	// ExpiryAfter Number of seconds after which this snapshot is considered expired by snapshot TTL.
	ExpiryAfter string `json:"expiry_after,omitempty"`
	// OriginName Origination group name.
	OriginName string `json:"origin_name,omitempty"`
	// IsReplica Snapshot is a replica from upstream replication partner.
	IsReplica bool `json:"is_replica,omitempty"`
	// IsUnmanaged Indicates whether the snapshot is unmanaged.
	IsUnmanaged bool `json:"is_unmanaged,omitempty"`
	// IsManuallyManaged Is snapshot manually managed, i.
	IsManuallyManaged bool `json:"is_manually_managed,omitempty"`
	// ReplicationStatus Replication status.
	ReplicationStatus string `json:"replication_status,omitempty"`
	// AccessControlRecords List of access control records that apply to this snapshot.
	AccessControlRecords []any `json:"access_control_records,omitempty"`
	// SerialNumber Identifier for the SCSI protocol.
	SerialNumber int64 `json:"serial_number,omitempty"`
	// TargetName The iSCSI Qualified Name (IQN) or the Fibre Channel World Wide Node Name (WWNN) of the target snapshot.
	TargetName string `json:"target_name,omitempty"`
	// CreationTime Time when this snapshot was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this snapshort was last modified.
	LastModified string `json:"last_modified,omitempty"`
	// ScheduleName Name of protection schedule.
	ScheduleName string `json:"schedule_name,omitempty"`
	// ScheduleID Identifier of protection schedule.
	ScheduleID string `json:"schedule_id,omitempty"`
	// AppUUID Application identifier of snapshots.
	AppUUID string `json:"app_uuid,omitempty"`
	// Metadata Key-value pairs that augment a snapshot's attributes.
	Metadata []any `json:"metadata,omitempty"`
	// NewDataValid Indicate the usage infomation is valid.
	NewDataValid string `json:"new_data_valid,omitempty"`
	// NewDataCompressedBytes The bytes of compressed new data.
	NewDataCompressedBytes int64 `json:"new_data_compressed_bytes,omitempty"`
	// NewDataUncompressedBytes The bytes of uncompressed new data.
	NewDataUncompressedBytes int64 `json:"new_data_uncompressed_bytes,omitempty"`
	// AgentType External management agent type.
	AgentType int64 `json:"agent_type,omitempty"`
	// VPDT10 The snapshot's T10 Vendor ID-based identifier.
	VPDT10 string `json:"vpd_t10,omitempty"`
	// VPDIeee0 The first 64 bits of the snapshots's EUI-64 identifier, encoded as a hexadecimal string.
	VPDIeee0 string `json:"vpd_ieee0,omitempty"`
	// VPDIeee1 The last 64 bits of the snapshots's EUI-64 identifier, encoded as a hexadecimal string.
	VPDIeee1 string `json:"vpd_ieee1,omitempty"`
	// Force Forcibly delete the specified snapshot even if it is the last replicated collection.
	Force bool `json:"force,omitempty"`
}

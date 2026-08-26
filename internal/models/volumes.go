package models

// Volume models the NimbleOS volumes object set.
type Volume struct {
	// ID Identifier for the volume.
	ID string `json:"id,omitempty"`
	// Name Name of the volume.
	Name string `json:"name,omitempty"`
	// FullName Fully qualified name of volume.
	FullName string `json:"full_name,omitempty"`
	// SearchName Name of volume used for object search.
	SearchName string `json:"search_name,omitempty"`
	// Size Volume size in mebibytes.
	Size uint64 `json:"size,omitempty"`
	// Description Text description of volume.
	Description string `json:"description,omitempty"`
	// PerfpolicyName Name of performance policy.
	PerfpolicyName string `json:"perfpolicy_name,omitempty"`
	// PerfpolicyID Identifier of the performance policy.
	PerfpolicyID string `json:"perfpolicy_id,omitempty"`
	// Reserve Amount of space to reserve for this volume as a percentage of volume size.
	Reserve uint64 `json:"reserve,omitempty"`
	// WarnLevel This attribute is deprecated.
	WarnLevel uint64 `json:"warn_level,omitempty"`
	// Limit Limit on the volume's mapped usage, expressed as a percentage of the volume's size.
	Limit uint64 `json:"limit,omitempty"`
	// SnapReserve Amount of space to reserve for snapshots of this volume as a percentage of volume size.
	SnapReserve uint64 `json:"snap_reserve,omitempty"`
	// SnapWarnLevel Threshold for available space as a percentage of volume size below which an alert is raised.
	SnapWarnLevel uint64 `json:"snap_warn_level,omitempty"`
	// SnapLimit This attribute is deprecated.
	SnapLimit uint64 `json:"snap_limit,omitempty"`
	// SnapLimitPercent This attribute is deprecated.
	SnapLimitPercent int64 `json:"snap_limit_percent,omitempty"`
	// NumSnaps Number of live, non-hidden snapshots for this volume.
	NumSnaps uint64 `json:"num_snaps,omitempty"`
	// ProjectedNumSnaps Deprecated.
	ProjectedNumSnaps uint64 `json:"projected_num_snaps,omitempty"`
	// Online Online state of volume, available for host initiators to establish connections.
	Online bool `json:"online,omitempty"`
	// OwnedByGroup Name of group that currently owns the volume.
	OwnedByGroup string `json:"owned_by_group,omitempty"`
	// OwnedByGroupID ID of group that currently owns the volume.
	OwnedByGroupID string `json:"owned_by_group_id,omitempty"`
	// MultiInitiator For iSCSI Volume Target, this flag indicates whether the volume and its snapshots can be accessed from multiple initiators at the same time.
	MultiInitiator bool `json:"multi_initiator,omitempty"`
	// PoolName Name of the pool where the volume resides.
	PoolName string `json:"pool_name,omitempty"`
	// PoolID Identifier associated with the pool in the storage pool table.
	PoolID string `json:"pool_id,omitempty"`
	// ReadOnly Volume is read-only.
	ReadOnly bool `json:"read_only,omitempty"`
	// SerialNumber Identifier associated with the volume for the SCSI protocol.
	SerialNumber string `json:"serial_number,omitempty"`
	// TargetName The iSCSI Qualified Name (IQN) or the Fibre Channel World Wide Node Name (WWNN) of the target volume.
	TargetName string `json:"target_name,omitempty"`
	// BlockSize Size in bytes of blocks in the volume.
	BlockSize uint64 `json:"block_size,omitempty"`
	// OfflineReason Volume offline reason.
	OfflineReason string `json:"offline_reason,omitempty"`
	// Clone Whether this volume is a clone.
	Clone bool `json:"clone,omitempty"`
	// ParentVolName Name of parent volume.
	ParentVolName string `json:"parent_vol_name,omitempty"`
	// ParentVolID Parent volume ID.
	ParentVolID string `json:"parent_vol_id,omitempty"`
	// BaseSnapName Name of base snapshot.
	BaseSnapName string `json:"base_snap_name,omitempty"`
	// BaseSnapID Base snapshot ID.
	BaseSnapID string `json:"base_snap_id,omitempty"`
	// VolcollName Name of volume collection of which this volume is a member.
	VolcollName string `json:"volcoll_name,omitempty"`
	// VolcollID ID of volume collection of which this volume is a member.
	VolcollID string `json:"volcoll_id,omitempty"`
	// AgentType External management agent type.
	AgentType string `json:"agent_type,omitempty"`
	// Force Forcibly offline, reduce size or change read-only status a volume.
	Force bool `json:"force,omitempty"`
	// CreationTime Time when this volume was created.
	CreationTime uint64 `json:"creation_time,omitempty"`
	// LastModified Time when this volume was last modified.
	LastModified uint64 `json:"last_modified,omitempty"`
	// ProtectionType Specifies if volume is protected with schedules.
	ProtectionType string `json:"protection_type,omitempty"`
	// LastSnap Last snapshot for this volume.
	LastSnap any `json:"last_snap,omitempty"`
	// LastReplicatedSnap Last replicated snapshot for this volume.
	LastReplicatedSnap any `json:"last_replicated_snap,omitempty"`
	// DestPoolName Name of the destination pool where the volume is moving to.
	DestPoolName string `json:"dest_pool_name,omitempty"`
	// DestPoolID ID of the destination pool where the volume is moving to.
	DestPoolID string `json:"dest_pool_id,omitempty"`
	// MoveStartTime The Start time when this volume was moved.
	MoveStartTime uint64 `json:"move_start_time,omitempty"`
	// MoveAborting This indicates whether the move of the volume is aborting or not.
	MoveAborting bool `json:"move_aborting,omitempty"`
	// MoveBytesMigrated The bytes of volume which have been moved.
	MoveBytesMigrated uint64 `json:"move_bytes_migrated,omitempty"`
	// MoveBytesRemaining The bytes of volume which have not been moved.
	MoveBytesRemaining uint64 `json:"move_bytes_remaining,omitempty"`
	// MoveEstComplTime The estimated time of completion of a move.
	MoveEstComplTime uint64 `json:"move_est_compl_time,omitempty"`
	// UsageValid This indicates whether usage information of volume and snapshots are valid or not.
	UsageValid bool `json:"usage_valid,omitempty"`
	// SpaceUsageLevel Indicates space usage level based on warning level.
	SpaceUsageLevel string `json:"space_usage_level,omitempty"`
	// TotalUsageBytes Sum of volume mapped usage and uncompressed backup data(including pending deletes) in bytes of this volume.
	TotalUsageBytes uint64 `json:"total_usage_bytes,omitempty"`
	// VolUsageCompressedBytes Compressed data in bytes for this volume.
	VolUsageCompressedBytes uint64 `json:"vol_usage_compressed_bytes,omitempty"`
	// VolUsageUncompressedBytes Uncompressed data in bytes for this volume.
	VolUsageUncompressedBytes uint64 `json:"vol_usage_uncompressed_bytes,omitempty"`
	// SnapUsageCompressedBytes Sum of compressed backup data in bytes stored in snapshots of this volume.
	SnapUsageCompressedBytes uint64 `json:"snap_usage_compressed_bytes,omitempty"`
	// SnapUsageUncompressedBytes Sum of uncompressed unique backup data in bytes stored in snapshots of this volume.
	SnapUsageUncompressedBytes uint64 `json:"snap_usage_uncompressed_bytes,omitempty"`
	// SnapUsagePopulatedBytes Sum of backup data in bytes stored in snapshots of this volume without accounting for the sharing of data between snapshots.
	SnapUsagePopulatedBytes uint64 `json:"snap_usage_populated_bytes,omitempty"`
	// CachePinned If set to true, all the contents of this volume are kept in flash cache.
	CachePinned bool `json:"cache_pinned,omitempty"`
	// PinnedCacheSize The amount of flash pinned on the volume.
	PinnedCacheSize uint64 `json:"pinned_cache_size,omitempty"`
	// CacheNeededForPin The amount of flash needed to pin the volume.
	CacheNeededForPin uint64 `json:"cache_needed_for_pin,omitempty"`
	// UpstreamCachePinned This indicates whether the upstream volume is cache pinned or not.
	UpstreamCachePinned bool `json:"upstream_cache_pinned,omitempty"`
	// CachePolicy Cache policy applied to the volume.
	CachePolicy string `json:"cache_policy,omitempty"`
	// ThinlyProvisioned Set volume's provisioning level to thin.
	ThinlyProvisioned bool `json:"thinly_provisioned,omitempty"`
	// VolState Status of the volume.
	VolState string `json:"vol_state,omitempty"`
	// OnlineSnaps The list of online snapshots of this volume.
	OnlineSnaps []any `json:"online_snaps,omitempty"`
	// NumConnections Number of connections of this volume.
	NumConnections uint64 `json:"num_connections,omitempty"`
	// NumISCSIConnections Number of iscsi connections of this volume.
	NumISCSIConnections uint64 `json:"num_iscsi_connections,omitempty"`
	// NumFCConnections Number of Fibre Channel connections of this volume.
	NumFCConnections uint64 `json:"num_fc_connections,omitempty"`
	// AccessControlRecords List of access control records that apply to this volume.
	AccessControlRecords []any `json:"access_control_records,omitempty"`
	// EncryptionCipher The encryption cipher of the volume.
	EncryptionCipher string `json:"encryption_cipher,omitempty"`
	// AppUUID Application identifier of volume.
	AppUUID string `json:"app_uuid,omitempty"`
	// FolderID ID of the folder holding this volume.
	FolderID string `json:"folder_id,omitempty"`
	// FolderName Name of the folder holding this volume.
	FolderName string `json:"folder_name,omitempty"`
	// Metadata Key-value pairs that augment an volume's attributes.
	Metadata []any `json:"metadata,omitempty"`
	// ISCSISessions List of iSCSI sessions connected to this volume.
	ISCSISessions []any `json:"iscsi_sessions,omitempty"`
	// FCSessions List of Fibre Channel sessions connected to this volume.
	FCSessions []any `json:"fc_sessions,omitempty"`
	// CachingEnabled Indicate caching the volume is enabled.
	CachingEnabled bool `json:"caching_enabled,omitempty"`
	// PreviouslyDeduped Indicate whether dedupe has ever been enabled on this volume.
	PreviouslyDeduped bool `json:"previously_deduped,omitempty"`
	// DedupeEnabled Indicate whether dedupe is enabled.
	DedupeEnabled bool `json:"dedupe_enabled,omitempty"`
	// VPDT10 The volume's T10 Vendor ID-based identifier.
	VPDT10 string `json:"vpd_t10,omitempty"`
	// VPDIeee0 The first 64 bits of the volume's EUI-64 identifier, encoded as a hexadecimal string.
	VPDIeee0 string `json:"vpd_ieee0,omitempty"`
	// VPDIeee1 The last 64 bits of the volume's EUI-64 identifier, encoded as a hexadecimal string.
	VPDIeee1 string `json:"vpd_ieee1,omitempty"`
	// AppCategory Application category that the volume belongs to.
	AppCategory string `json:"app_category,omitempty"`
	// LimitIOPS IOPS limit for this volume.
	LimitIOPS int64 `json:"limit_iops,omitempty"`
	// LimitMBPS Throughput limit for this volume in MB/s.
	LimitMBPS int64 `json:"limit_mbps,omitempty"`
	// PreFilter Pre-filtering criteria.
	PreFilter map[string]any `json:"pre_filter,omitempty"`
	// AvgStatsLast5mins Average statistics in last 5 minutes.
	AvgStatsLast5mins VolumeStats `json:"avg_stats_last_5mins"`
}

type VolumeStats struct {
	ReadIops           uint64 `json:"read_iops,omitempty"`
	ReadThroughput     uint64 `json:"read_throughput,omitempty"`
	ReadLatency        uint64 `json:"read_latency,omitempty"`
	WriteIops          uint64 `json:"write_iops,omitempty"`
	WriteThroughput    uint64 `json:"write_throughput,omitempty"`
	WriteLatency       uint64 `json:"write_latency,omitempty"`
	CombinedIops       uint64 `json:"combined_iops,omitempty"`
	CombinedThroughput uint64 `json:"combined_throughput,omitempty"`
	CombinedLatency    uint64 `json:"combined_latency,omitempty"`
}

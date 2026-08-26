package models

// Folder models the NimbleOS folders object set.
type Folder struct {
	// ID Identifier for the folder.
	ID string `json:"id,omitempty"`
	// Name Name of folder.
	Name string `json:"name,omitempty"`
	// Fqn Fully qualified name of folder in the pool.
	Fqn string `json:"fqn,omitempty"`
	// FullName Fully qualified name of folder in the group.
	FullName string `json:"full_name,omitempty"`
	// SearchName Name of folder used for object search.
	SearchName string `json:"search_name,omitempty"`
	// Description Text description of folder.
	Description string `json:"description,omitempty"`
	// PoolName Name of the pool where the folder resides.
	PoolName string `json:"pool_name,omitempty"`
	// PoolID ID of the pool where the folder resides.
	PoolID string `json:"pool_id,omitempty"`
	// LimitBytesSpecified Indicates whether the folder has a limit.
	LimitBytesSpecified bool `json:"limit_bytes_specified,omitempty"`
	// LimitBytes Folder limit size in bytes.
	LimitBytes int64 `json:"limit_bytes,omitempty"`
	// LimitSizeBytes Folder size limit in bytes.
	LimitSizeBytes int64 `json:"limit_size_bytes,omitempty"`
	// ProvisionedLimitSizeBytes Limit on the provisioned size of volumes in a folder.
	ProvisionedLimitSizeBytes int64 `json:"provisioned_limit_size_bytes,omitempty"`
	// OverdraftLimitPct Amount of space to consider as overdraft range for this folder as a percentage of folder used limit.
	OverdraftLimitPct int64 `json:"overdraft_limit_pct,omitempty"`
	// CapacityBytes Capacity of the folder in bytes.
	CapacityBytes int64 `json:"capacity_bytes,omitempty"`
	// FreeSpaceBytes Free space in the folder in bytes.
	FreeSpaceBytes int64 `json:"free_space_bytes,omitempty"`
	// ProvisionedBytes Sum of provisioned size of volumes in the folder.
	ProvisionedBytes int64 `json:"provisioned_bytes,omitempty"`
	// UsageBytes Sum of mapped usage and snapshot uncompressed usage of volumes in the folder.
	UsageBytes int64 `json:"usage_bytes,omitempty"`
	// VolumeMappedBytes Sum of mapped usage of volumes in the folder.
	VolumeMappedBytes int64 `json:"volume_mapped_bytes,omitempty"`
	// UsageValid Indicate whether the space usage attributes of folder are valid.
	UsageValid bool `json:"usage_valid,omitempty"`
	// AgentType External management agent type.
	AgentType int64 `json:"agent_type,omitempty"`
	// InheritedVolPerfpolID Identifier of the default performance policy for a newly created volume.
	InheritedVolPerfpolID string `json:"inherited_vol_perfpol_id,omitempty"`
	// InheritedVolPerfpolName Name of the default performance policy for a newly created volume.
	InheritedVolPerfpolName string `json:"inherited_vol_perfpol_name,omitempty"`
	// UnusedReserveBytes Unused reserve of volumes in the folder in bytes.
	UnusedReserveBytes int64 `json:"unused_reserve_bytes,omitempty"`
	// UnusedSnapReserveBytes Unused reserve of snapshots of volumes in the folder in bytes.
	UnusedSnapReserveBytes int64 `json:"unused_snap_reserve_bytes,omitempty"`
	// CompressedVolUsageBytes Compressed usage of volumes in the folder.
	CompressedVolUsageBytes int64 `json:"compressed_vol_usage_bytes,omitempty"`
	// CompressedSnapUsageBytes Compressed usage of snapshots in the folder.
	CompressedSnapUsageBytes int64 `json:"compressed_snap_usage_bytes,omitempty"`
	// UncompressedVolUsageBytes Uncompressed usage of volumes in the folder.
	UncompressedVolUsageBytes int64 `json:"uncompressed_vol_usage_bytes,omitempty"`
	// UncompressedSnapUsageBytes Uncompressed usage of snapshots in the folder.
	UncompressedSnapUsageBytes int64 `json:"uncompressed_snap_usage_bytes,omitempty"`
	// VolCompressionRatio Compression ratio of volumes in the folder.
	VolCompressionRatio float64 `json:"vol_compression_ratio,omitempty"`
	// SnapCompressionRatio Compression ratio of snapshots in the folder.
	SnapCompressionRatio float64 `json:"snap_compression_ratio,omitempty"`
	// CompressionRatio Compression savings for the folder expressed as ratio.
	CompressionRatio float64 `json:"compression_ratio,omitempty"`
	// CreationTime Time when this folder was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this folder was last modified.
	LastModified string `json:"last_modified,omitempty"`
	// NumSnaps Number of snapshots inside the folder.
	NumSnaps string `json:"num_snaps,omitempty"`
	// NumSnapcolls Number of snapshot collections inside the folder.
	NumSnapcolls string `json:"num_snapcolls,omitempty"`
	// AppUUID Application identifier of the folder.
	AppUUID string `json:"app_uuid,omitempty"`
	// VolumeList List of volumes contained by the folder.
	VolumeList []any `json:"volume_list,omitempty"`
	// AppserverID Identifier of the application server associated with the folder.
	AppserverID string `json:"appserver_id,omitempty"`
	// AppserverName Name of the application server associated with the folder.
	AppserverName string `json:"appserver_name,omitempty"`
	// FolsetID Identifier of the folder set associated with the folder.
	FolsetID string `json:"folset_id,omitempty"`
	// FolsetName Name of the folder set associated with the folder.
	FolsetName string `json:"folset_name,omitempty"`
	// LimitIOPS IOPS limit for this folder.
	LimitIOPS int64 `json:"limit_iops,omitempty"`
	// LimitMBPS Throughput limit for this folder in MB/s.
	LimitMBPS int64 `json:"limit_mbps,omitempty"`
	// AccessProtocol Access protocol of the folder.
	AccessProtocol string `json:"access_protocol,omitempty"`
	// TenantID Tenant ID of the folder.
	TenantID string `json:"tenant_id,omitempty"`
}

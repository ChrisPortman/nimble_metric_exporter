package models

// Array models the NimbleOS arrays object set.
type Array struct {
	// ID Identifier for array.
	ID string `json:"id,omitempty"`
	// Name The user provided name of the array.
	Name string `json:"name,omitempty"`
	// Force Forcibly delete the specified array.
	Force bool `json:"force,omitempty"`
	// FullName The array's fully qualified name.
	FullName string `json:"full_name,omitempty"`
	// SearchName The array name used for object search.
	SearchName string `json:"search_name,omitempty"`
	// Status Reachability status of the array in the group.
	Status string `json:"status,omitempty"`
	// Role Role of the array in the group.
	Role string `json:"role,omitempty"`
	// GroupState State of the array in the group.
	GroupState string `json:"group_state,omitempty"`
	// PoolName Name of pool to which this is a member.
	PoolName string `json:"pool_name,omitempty"`
	// PoolID ID of pool to which this is a member.
	PoolID string `json:"pool_id,omitempty"`
	// Model Array model.
	Model string `json:"model,omitempty"`
	// Serial Serial number of the array.
	Serial string `json:"serial,omitempty"`
	// Version Software version of the array.
	Version string `json:"version,omitempty"`
	// IsSfa True if this array supports SFA; false otherwise.
	IsSfa bool `json:"is_sfa,omitempty"`
	// CreationTime Time when this array object was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this array object was last modified.
	LastModified string `json:"last_modified,omitempty"`
	// UsageValid Indicates whether the usage of array is valid.
	UsageValid bool `json:"usage_valid,omitempty"`
	// UsableCapacityBytes The usable capacity of the array in bytes.
	UsableCapacityBytes int64 `json:"usable_capacity_bytes,omitempty"`
	// UsableCacheCapacityBytes The usable cache capacity of the array in bytes.
	UsableCacheCapacityBytes int64 `json:"usable_cache_capacity_bytes,omitempty"`
	// RawCapacityBytes The raw capacity of the array in bytes.
	RawCapacityBytes int64 `json:"raw_capacity_bytes,omitempty"`
	// VolUsageBytes The compressed usage of volumes in array.
	VolUsageBytes int64 `json:"vol_usage_bytes,omitempty"`
	// VolUsageUncompressedBytes The uncompressed usage of volumes in array.
	VolUsageUncompressedBytes int64 `json:"vol_usage_uncompressed_bytes,omitempty"`
	// VolCompression The compression rate of volumes in array expressed as ratio.
	VolCompression float64 `json:"vol_compression,omitempty"`
	// VolSavedBytes The saved space of volumes in array.
	VolSavedBytes int64 `json:"vol_saved_bytes,omitempty"`
	// SnapUsageBytes The compressed usage of snapshots in array.
	SnapUsageBytes int64 `json:"snap_usage_bytes,omitempty"`
	// SnapUsageUncompressedBytes The uncompressed usage of snapshots in array.
	SnapUsageUncompressedBytes int64 `json:"snap_usage_uncompressed_bytes,omitempty"`
	// SnapCompression The compression rate of snapshots in array expressed as ratio.
	SnapCompression float64 `json:"snap_compression,omitempty"`
	// SnapSpaceReduction The space reduction rate of snapshots in array expressed as ratio.
	SnapSpaceReduction float64 `json:"snap_space_reduction,omitempty"`
	// SnapSavedBytes The saved space of snapshots in array.
	SnapSavedBytes int64 `json:"snap_saved_bytes,omitempty"`
	// PendingDeleteBytes The pending delete bytes in array.
	PendingDeleteBytes int64 `json:"pending_delete_bytes,omitempty"`
	// AvailableBytes The available space of array.
	AvailableBytes int64 `json:"available_bytes,omitempty"`
	// Usage Used space of the array in bytes.
	Usage int64 `json:"usage,omitempty"`
	// AllFlash Whether it is an all-flash array.
	AllFlash bool `json:"all_flash,omitempty"`
	// DedupeCapacityBytes The dedupe capacity of a hybrid array.
	DedupeCapacityBytes int64 `json:"dedupe_capacity_bytes,omitempty"`
	// DedupeUsageBytes The dedupe usage of a hybrid array.
	DedupeUsageBytes int64 `json:"dedupe_usage_bytes,omitempty"`
	// IsFullyDedupeCapable Is array fully capable to dedupe its usable capacity.
	IsFullyDedupeCapable bool `json:"is_fully_dedupe_capable,omitempty"`
	// DedupeDisabled Is data deduplication disabled for this array.
	DedupeDisabled string `json:"dedupe_disabled,omitempty"`
	// ExtendedModel Extended model of the array.
	ExtendedModel string `json:"extended_model,omitempty"`
	// Oem OEM brand of the array.
	Oem string `json:"oem,omitempty"`
	// Brand Brand of the array.
	Brand string `json:"brand,omitempty"`
	// IsSupportedHwConfig Whether it is a supported hardware config.
	IsSupportedHwConfig bool `json:"is_supported_hw_config,omitempty"`
	// GigNICPortCount Count of 1G NIC Ports installed on the array.
	GigNICPortCount int64 `json:"gig_nic_port_count,omitempty"`
	// TenGigSfpNICPortCount Count of SFP NIC Ports installed on the array capable of 10G, 25G or 100G speeds.
	TenGigSfpNICPortCount int64 `json:"ten_gig_sfp_nic_port_count,omitempty"`
	// TenGigTNICPortCount Count of 10G BaseT NIC Ports installed on the array.
	TenGigTNICPortCount int64 `json:"ten_gig_t_nic_port_count,omitempty"`
	// FCPortCount Count of Fibre Channel Ports installed on the array.
	FCPortCount int64 `json:"fc_port_count,omitempty"`
	// PublicKey Public key of the array.
	PublicKey string `json:"public_key,omitempty"`
	// Upgrade The array upgrade data.
	Upgrade map[string]any `json:"upgrade,omitempty"`
	// CreatePool Whether to create associated pool during array create.
	CreatePool bool `json:"create_pool,omitempty"`
	// PoolDescription Text description of the pool to be created during array creation.
	PoolDescription string `json:"pool_description,omitempty"`
	// AllowLowerLimits A True setting will allow you to add an array with lower limits to a pool with higher limits.
	AllowLowerLimits bool `json:"allow_lower_limits,omitempty"`
	// CtrlrASupportIP Controller A Support IP Address.
	CtrlrASupportIP int64 `json:"ctrlr_a_support_ip,omitempty"`
	// CtrlrBSupportIP Controller B Support IP Address.
	CtrlrBSupportIP int64 `json:"ctrlr_b_support_ip,omitempty"`
	// NICList List NICs information.
	NICList []any `json:"nic_list,omitempty"`
	// ModelSubType Array model sub type.
	ModelSubType string `json:"model_sub_type,omitempty"`
	// ZconfIpaddrs List of link-local zero-configuration addresses of the array.
	ZconfIpaddrs []any `json:"zconf_ipaddrs,omitempty"`
	// SecondaryMgmtIP Secondary management IP address for the Group.
	SecondaryMgmtIP string `json:"secondary_mgmt_ip,omitempty"`
}

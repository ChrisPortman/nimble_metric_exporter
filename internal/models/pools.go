package models

// Pool models the NimbleOS pools object set.
type Pool struct {
	// ID Identifier for the pool.
	ID string `json:"id,omitempty"`
	// Name Name of pool.
	Name string `json:"name,omitempty"`
	// FullName Fully qualified name of pool.
	FullName string `json:"full_name,omitempty"`
	// SearchName Name of pool used for object search.
	SearchName string `json:"search_name,omitempty"`
	// Description Text description of pool.
	Description string `json:"description,omitempty"`
	// CreationTime Time when this pool was created.
	CreationTime uint64 `json:"creation_time,omitempty"`
	// LastModified Time when this pool was last modified.
	LastModified uint64 `json:"last_modified,omitempty"`
	// Capacity Total storage space of the pool in bytes.
	Capacity uint64 `json:"capacity,omitempty"`
	// Usage Used space of the pool in bytes.
	Usage uint64 `json:"usage,omitempty"`
	// Savings Overall space usage savings in the pool.
	Savings uint64 `json:"savings,omitempty"`
	// SavingsDataReduction Space usage savings in the pool that does not include thin-provisioning savings.
	SavingsDataReduction uint64 `json:"savings_data_reduction,omitempty"`
	// SavingsCompression Space usage savings in the pool due to compression.
	SavingsCompression uint64 `json:"savings_compression,omitempty"`
	// SavingsDedupe Space usage savings in the pool due to deduplication.
	SavingsDedupe uint64 `json:"savings_dedupe,omitempty"`
	// SavingsClone Space usage savings in the pool due to cloning of volumes.
	SavingsClone uint64 `json:"savings_clone,omitempty"`
	// SavingsVolThinProvisioning Space usage savings in the pool due to thin provisioning of volumes.
	SavingsVolThinProvisioning uint64 `json:"savings_vol_thin_provisioning,omitempty"`
	// Reserve Reserved space of the pool in bytes.
	Reserve uint64 `json:"reserve,omitempty"`
	// UnusedReserve Unused reserve space of the pool in bytes.
	UnusedReserve uint64 `json:"unused_reserve,omitempty"`
	// FreeSpace Free space of the pool in bytes.
	FreeSpace uint64 `json:"free_space,omitempty"`
	// CacheCapacity Total usable cache capacity of the pool in bytes.
	CacheCapacity uint64 `json:"cache_capacity,omitempty"`
	// PinnableCacheCapacity Total pinnable cache capacity of the pool in bytes.
	PinnableCacheCapacity uint64 `json:"pinnable_cache_capacity,omitempty"`
	// PinnedCacheCapacity Total pinned cache capacity of the pool in bytes.
	PinnedCacheCapacity uint64 `json:"pinned_cache_capacity,omitempty"`
	// DedupeCapacityBytes The dedupe capacity of a hybrid pool.
	DedupeCapacityBytes uint64 `json:"dedupe_capacity_bytes,omitempty"`
	// DedupeUsageBytes The dedupe usage of a hybrid pool.
	DedupeUsageBytes uint64 `json:"dedupe_usage_bytes,omitempty"`
	// SavingsRatio Overall space usage savings in the pool expressed as ratio.
	SavingsRatio float64 `json:"savings_ratio,omitempty"`
	// DataReductionRatio Space usage savings in the pool expressed as ratio that does not include thin-provisioning savings.
	DataReductionRatio float64 `json:"data_reduction_ratio,omitempty"`
	// CompressionRatio Compression savings for the pool expressed as ratio.
	CompressionRatio float64 `json:"compression_ratio,omitempty"`
	// DedupeRatio Dedupe savings for the pool expressed as ratio.
	DedupeRatio float64 `json:"dedupe_ratio,omitempty"`
	// CloneRatio Clone savings for the pool expressed as ratio.
	CloneRatio float64 `json:"clone_ratio,omitempty"`
	// VolThinProvisioningRatio Thin provisioning savings for volumes in the pool expressed as ratio.
	VolThinProvisioningRatio float64 `json:"vol_thin_provisioning_ratio,omitempty"`
	// SnapcollCount Snapshot collection count.
	SnapcollCount uint64 `json:"snapcoll_count,omitempty"`
	// SnapCount Snapshot count.
	SnapCount uint64 `json:"snap_count,omitempty"`
	// ArrayCount Number of arrays in the pool.
	ArrayCount uint64 `json:"array_count,omitempty"`
	// VolCount Number of volumes assigned to the pool.
	VolCount uint64 `json:"vol_count,omitempty"`
	// ArrayList List of arrays in the pool with detailed information.
	ArrayList []any `json:"array_list,omitempty"`
	// UnassignedArrayList List of arrays being unassigned from the pool.
	UnassignedArrayList []any `json:"unassigned_array_list,omitempty"`
	// VolList The list of volumes in the pool.
	VolList []any `json:"vol_list,omitempty"`
	// PinnedVolList The list of pinned volumes in the pool.
	PinnedVolList []any `json:"pinned_vol_list,omitempty"`
	// FolderList The list of fully qualified names of folders in the pool.
	FolderList []any `json:"folder_list,omitempty"`
	// Force Forcibly delete the specified pool even if it contains deleted volumes whose space is being reclaimed.
	Force bool `json:"force,omitempty"`
	// UsageValid Indicates whether the usage of pool is valid.
	UsageValid bool `json:"usage_valid,omitempty"`
	// UncompressedVolUsageBytes Uncompressed usage of volumes in the pool.
	UncompressedVolUsageBytes int64 `json:"uncompressed_vol_usage_bytes,omitempty"`
	// UncompressedSnapUsageBytes Uncompressed usage of snapshots in the pool.
	UncompressedSnapUsageBytes int64 `json:"uncompressed_snap_usage_bytes,omitempty"`
	// AllFlash Indicate whether the pool is an all_flash pool.
	AllFlash bool `json:"all_flash,omitempty"`
	// DedupeCapable Indicates whether the pool is capable of hosting deduped volumes.
	DedupeCapable bool `json:"dedupe_capable,omitempty"`
	// DedupeAllVolumesCapable Indicates whether the pool can enable dedupe by default.
	DedupeAllVolumesCapable bool `json:"dedupe_all_volumes_capable,omitempty"`
	// DedupeAllVolumes Indicates if dedupe is enabled by default for new volumes on this pool.
	DedupeAllVolumes bool `json:"dedupe_all_volumes,omitempty"`
	// IsDefault Indicates if this is the default pool.
	IsDefault bool `json:"is_default,omitempty"`
}

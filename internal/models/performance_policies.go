package models

// PerformancePolicy models the NimbleOS performance_policies object set.
type PerformancePolicy struct {
	// ID Unique Identifier for the Performance Policy.
	ID string `json:"id,omitempty"`
	// Name Name of the Performance Policy.
	Name string `json:"name,omitempty"`
	// FullName Fully qualified name of the Performance Policy.
	FullName string `json:"full_name,omitempty"`
	// SearchName Name of the Performance Policy used for object search.
	SearchName string `json:"search_name,omitempty"`
	// Description Description of a performance policy.
	Description string `json:"description,omitempty"`
	// BlockSize Block Size in bytes to be used by the volumes created with this specific performance policy.
	BlockSize int64 `json:"block_size,omitempty"`
	// Compress Flag denoting if data in the associated volume should be compressed.
	Compress bool `json:"compress,omitempty"`
	// Cache Flag denoting if data in the associated volume should be cached.
	Cache bool `json:"cache,omitempty"`
	// CachePolicy Specifies how data of associated volume should be cached.
	CachePolicy string `json:"cache_policy,omitempty"`
	// SpacePolicy Specifies the state of the volume upon space constraint violation such as volume limit violation or volumes above their volume reserve, if the pool free space is exhausted.
	SpacePolicy int64 `json:"space_policy,omitempty"`
	// AppCategory Specifies the application category of the associated volume.
	AppCategory string `json:"app_category,omitempty"`
	// DedupeEnabled Specifies if dedupe is enabled for volumes created with this performance policy.
	DedupeEnabled bool `json:"dedupe_enabled,omitempty"`
	// Deprecated Specifies if this performance policy is deprecated.
	Deprecated bool `json:"deprecated,omitempty"`
	// Predefined Specifies if this performance policy is predefined (read-only).
	Predefined bool `json:"predefined,omitempty"`
	// CreationTime Time when the performance policy was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when the performance policy's configurations were last modified.
	LastModified string `json:"last_modified,omitempty"`
	// SampleRate Sample rate value.
	SampleRate int64 `json:"sample_rate,omitempty"`
	// VolumeCount Number of volumes using this performance policy.
	VolumeCount int64 `json:"volume_count,omitempty"`
	// DedupeOverridePools List of pools that override performance policy's dedupe setting.
	DedupeOverridePools []any `json:"dedupe_override_pools,omitempty"`
}

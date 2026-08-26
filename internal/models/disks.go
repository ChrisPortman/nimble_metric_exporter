package models

// Disk models the NimbleOS disks object set.
type Disk struct {
	// ID ID of disk.
	ID string `json:"id,omitempty"`
	// IsDfc Is disk part of dual flash carrier.
	IsDfc bool `json:"is_dfc,omitempty"`
	// Serial Disk serial number(N/A if empty).
	Serial string `json:"serial,omitempty"`
	// Path Disk SCSI device path.
	Path string `json:"path,omitempty"`
	// ShelfSerial Serial number of the shelf the disk is attached to.
	ShelfSerial string `json:"shelf_serial,omitempty"`
	// ShelfLocation Identifies the controller, port, and chain position of the shelf the disk belongs to.
	ShelfLocation string `json:"shelf_location,omitempty"`
	// ShelfID Identifies the physical shelf the disk belongs to.
	ShelfID string `json:"shelf_id,omitempty"`
	// ShelfLocationID Identifies the position shelf the disk belongs to, as coded integer.
	ShelfLocationID int64 `json:"shelf_location_id,omitempty"`
	// VshelfID Identifies the local shelf ID the disk belongs to.
	VshelfID int64 `json:"vshelf_id,omitempty"`
	// Slot Disk slot number.
	Slot int64 `json:"slot,omitempty"`
	// Bank Disk bank number.
	Bank int64 `json:"bank,omitempty"`
	// Model Disk model name.
	Model string `json:"model,omitempty"`
	// Vendor Vendor name of the disk manufacturer.
	Vendor string `json:"vendor,omitempty"`
	// FirmwareVersion Firmware version on the disk.
	FirmwareVersion string `json:"firmware_version,omitempty"`
	// Hba HBA ID the disk is connected to.
	Hba int64 `json:"hba,omitempty"`
	// Port HBA port number the disk is connected to.
	Port int64 `json:"port,omitempty"`
	// Size Disk size in bytes.
	Size int64 `json:"size,omitempty"`
	// State Disk hardware state. One of 'valid', 'in use', 'failed', absent', 'removed', 'void', 't_fail', 'foreign'
	State string `json:"state,omitempty"`
	// Type Type of disk (HDD, SSD, N/A).
	Type string `json:"type,omitempty"`
	// BlockType Native block type of the disk.
	BlockType string `json:"block_type,omitempty"`
	// RaidID Raid ID.
	RaidID int64 `json:"raid_id,omitempty"`
	// RaidResyncPercent Percentage RAID rebuild completed on this disk.
	RaidResyncPercent int64 `json:"raid_resync_percent,omitempty"`
	// RaidResyncCurrentSpeed Current RAID rebuild speed (bytes/sec).
	RaidResyncCurrentSpeed int64 `json:"raid_resync_current_speed,omitempty"`
	// RaidResyncAverageSpeed Average RAID rebuild speed (bytes/sec).
	RaidResyncAverageSpeed int64 `json:"raid_resync_average_speed,omitempty"`
	// RaidState RAID status for the disk (N/A, okay, resynchronizing, spare, faulty).
	RaidState string `json:"raid_state,omitempty"`
	// DiskInternalStat1 Internal disk statistic 1.
	DiskInternalStat1 string `json:"disk_internal_stat_1,omitempty"`
	// SmartAttributeList S.
	SmartAttributeList []any `json:"smart_attribute_list,omitempty"`
	// DiskOp The intended operation to be performed on the specified disk.
	DiskOp string `json:"disk_op,omitempty"`
	// Force Forcibly add a disk.
	Force bool `json:"force,omitempty"`
	// ArrayName Name of array the disk belongs to.
	ArrayName string `json:"array_name,omitempty"`
	// ArrayID ID of array the disk belongs to.
	ArrayID string `json:"array_id,omitempty"`
	// PartialResponseOk Indicate that it is okay to provide partially available response.
	PartialResponseOk string `json:"partial_response_ok,omitempty"`
}

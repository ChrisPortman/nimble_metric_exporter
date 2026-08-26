package models

// ReplicationPartner models the NimbleOS replication_partners object set.
type ReplicationPartner struct {
	// ID Identifier for a replication partner.
	ID string `json:"id,omitempty"`
	// Name Name of replication partner.
	Name string `json:"name,omitempty"`
	// FullName Fully qualified name of replication partner.
	FullName string `json:"full_name,omitempty"`
	// SearchName Name of replication partner used for object search.
	SearchName string `json:"search_name,omitempty"`
	// Description Description of replication partner.
	Description string `json:"description,omitempty"`
	// PartnerType Replication partner type.
	PartnerType string `json:"partner_type,omitempty"`
	// Alias Name this group uses to identify itself to this partner.
	Alias string `json:"alias,omitempty"`
	// Secret Replication partner shared secret, used for mutual authentication of the partners.
	Secret string `json:"secret,omitempty"`
	// CreationTime Time when this replication partner was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this replication partner was last modified.
	LastModified string `json:"last_modified,omitempty"`
	// ControlPort Port number of partner control interface.
	ControlPort int64 `json:"control_port,omitempty"`
	// Hostname IP address or hostname of partner interface.
	Hostname string `json:"hostname,omitempty"`
	// PortRangeStart For tunnel_endpoint partner types, first port available on the ssh proxy available for reverse forwarding.
	PortRangeStart int64 `json:"port_range_start,omitempty"`
	// ProxyHostname IP address of tunnel endpoint.
	ProxyHostname string `json:"proxy_hostname,omitempty"`
	// ProxyUser User to authenticate with tunnel endpoint.
	ProxyUser string `json:"proxy_user,omitempty"`
	// ReplHostname IP address or hostname of partner data interface.
	ReplHostname string `json:"repl_hostname,omitempty"`
	// DataPort Port number of partner data interface.
	DataPort int64 `json:"data_port,omitempty"`
	// IsAlive Whether the partner is available, and responding to pings.
	IsAlive bool `json:"is_alive,omitempty"`
	// PartnerGroupUid Replication partner group uid.
	PartnerGroupUid string `json:"partner_group_uid,omitempty"`
	// LastKeepaliveError Most recent error while attempting to ping the partner.
	LastKeepaliveError string `json:"last_keepalive_error,omitempty"`
	// CfgSyncStatus Indicates whether all volumes and volume collections have been synced to the partner.
	CfgSyncStatus bool `json:"cfg_sync_status,omitempty"`
	// LastSyncError Most recent error seen while attempting to sync objects to the partner.
	LastSyncError string `json:"last_sync_error,omitempty"`
	// ArraySerial Serial number of group leader array of the partner.
	ArraySerial string `json:"array_serial,omitempty"`
	// Version Replication version of the partner.
	Version string `json:"version,omitempty"`
	// PoolID The pool ID where volumes replicated from this partner will be created.
	PoolID string `json:"pool_id,omitempty"`
	// PoolName The pool name where volumes replicated from this partner will be created.
	PoolName string `json:"pool_name,omitempty"`
	// FolderID The Folder ID within the pool where volumes replicated from this partner will be created.
	FolderID string `json:"folder_id,omitempty"`
	// FolderName The Folder name within the pool where volumes replicated from this partner will be created.
	FolderName string `json:"folder_name,omitempty"`
	// MatchFolder Indicates whether to match the upstream volume's folder on the downstream.
	MatchFolder bool `json:"match_folder,omitempty"`
	// Paused Indicates whether replication traffic from/to this partner has been halted.
	Paused bool `json:"paused,omitempty"`
	// UniqueName Indicates whether this partner actively mangles object names to avoid name conflicts during replication.
	UniqueName bool `json:"unique_name,omitempty"`
	// SubnetLabel Label of the subnet used to replicate to this partner.
	SubnetLabel string `json:"subnet_label,omitempty"`
	// SubnetType Type of the subnet used to replicate to this partner.
	SubnetType string `json:"subnet_type,omitempty"`
	// Throttles Throttles used while replicating from/to this partner.
	Throttles []any `json:"throttles,omitempty"`
	// ThrottledBandwidth Current bandwidth throttle for this partner, expressed either as megabits per second or as the largest possible 64-bit signed integer (9223372036854775807) to indicate that there is no throttle.
	ThrottledBandwidth int64 `json:"throttled_bandwidth,omitempty"`
	// ThrottledBandwidthCurrent Current bandwidth throttle for this partner, expressed either as megabits per second or as -1 to indicate that there is no throttle.
	ThrottledBandwidthCurrent int64 `json:"throttled_bandwidth_current,omitempty"`
	// ThrottledBandwidthKbps Current bandwidth throttle for this partner, expressed either as kilobits per second or as the largest possible 64-bit signed integer (9223372036854775807) to indicate that there is no throttle.
	ThrottledBandwidthKbps int64 `json:"throttled_bandwidth_kbps,omitempty"`
	// ThrottledBandwidthCurrentKbps Current bandwidth throttle for this partner, expressed either as kilobits per second or as -1 to indicate that there is no throttle.
	ThrottledBandwidthCurrentKbps int64 `json:"throttled_bandwidth_current_kbps,omitempty"`
	// SubnetNetwork Subnet used to replicate to this partner.
	SubnetNetwork string `json:"subnet_network,omitempty"`
	// SubnetNetmask Subnet mask used to replicate to this partner.
	SubnetNetmask string `json:"subnet_netmask,omitempty"`
	// VolumeCollectionList List of volume collections that are replicating from/to this partner.
	VolumeCollectionList []any `json:"volume_collection_list,omitempty"`
	// VolumeCollectionListCount Count of volume collections that are replicating from/to this partner.
	VolumeCollectionListCount []any `json:"volume_collection_list_count,omitempty"`
	// VolumeList List of volumes that are replicating from/to this partner.
	VolumeList []any `json:"volume_list,omitempty"`
	// VolumeListCount Count of volumes that are replicating from/to this partner.
	VolumeListCount []any `json:"volume_list_count,omitempty"`
	// ReplicationDirection Direction of replication configured with this partner.
	ReplicationDirection string `json:"replication_direction,omitempty"`
}

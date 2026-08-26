package models

// VolumeCollection models the NimbleOS volume_collections object set.
type VolumeCollection struct {
	// ID Identifier for volume collection.
	ID string `json:"id,omitempty"`
	// ProttmplID Identifier of the protection template whose attributes will be used to create this volume collection.
	ProttmplID string `json:"prottmpl_id,omitempty"`
	// Name Name of volume collection.
	Name string `json:"name,omitempty"`
	// FullName Fully qualified name of volume collection.
	FullName string `json:"full_name,omitempty"`
	// SearchName Name of volume collection used for object search.
	SearchName string `json:"search_name,omitempty"`
	// Description Text description of volume collection.
	Description string `json:"description,omitempty"`
	// ReplPriority Replication priority for the volume collection with the following choices: {normal | high}.
	ReplPriority int64 `json:"repl_priority,omitempty"`
	// PolOwnerName Owner group.
	PolOwnerName string `json:"pol_owner_name,omitempty"`
	// ReplicationType Type of replication configured for the volume collection.
	ReplicationType string `json:"replication_type,omitempty"`
	// SynchronousReplicationType Type of synchronous replication configured for the volume collection.
	SynchronousReplicationType string `json:"synchronous_replication_type,omitempty"`
	// SynchronousReplicationState State of synchronous replication on the volume collection.
	SynchronousReplicationState string `json:"synchronous_replication_state,omitempty"`
	// AppSync Application Synchronization.
	AppSync string `json:"app_sync,omitempty"`
	// AppServer Application server hostname.
	AppServer string `json:"app_server,omitempty"`
	// AppID Application ID running on the server.
	AppID string `json:"app_id,omitempty"`
	// AppClusterName If the application is running within a Windows cluster environment, this is the cluster name.
	AppClusterName string `json:"app_cluster_name,omitempty"`
	// AppServiceName If the application is running within a Windows cluster environment then this is the instance name of the service running within the cluster environment.
	AppServiceName string `json:"app_service_name,omitempty"`
	// VcenterHostname VMware vCenter hostname.
	VcenterHostname string `json:"vcenter_hostname,omitempty"`
	// VcenterUsername Application VMware vCenter username.
	VcenterUsername string `json:"vcenter_username,omitempty"`
	// VcenterPassword Application VMware vCenter password.
	VcenterPassword string `json:"vcenter_password,omitempty"`
	// AgentHostname Generic backup agent hostname.
	AgentHostname int64 `json:"agent_hostname,omitempty"`
	// AgentUsername Generic backup agent username.
	AgentUsername int64 `json:"agent_username,omitempty"`
	// AgentPassword Generic backup agent password.
	AgentPassword int64 `json:"agent_password,omitempty"`
	// CreationTime Time when this volume collection was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModifiedTime Time when this volume collection was last modified.
	LastModifiedTime int64 `json:"last_modified_time,omitempty"`
	// VolumeList List of volumes associated with the volume collection.
	VolumeList []any `json:"volume_list,omitempty"`
	// DownstreamVolumeList List of downstream volumes associated with the volume collection.
	DownstreamVolumeList []any `json:"downstream_volume_list,omitempty"`
	// UpstreamVolumeList List of upstream volumes associated with the volume collection.
	UpstreamVolumeList []any `json:"upstream_volume_list,omitempty"`
	// VolumeCount Count of volumes associated with the volume collection.
	VolumeCount int64 `json:"volume_count,omitempty"`
	// CachePinnedVolumeList List of cache pinned volumes associated with volume collection.
	CachePinnedVolumeList []any `json:"cache_pinned_volume_list,omitempty"`
	// LastSnapcoll Last snapshot collection on this volume collection.
	LastSnapcoll string `json:"last_snapcoll,omitempty"`
	// SnapcollCount Count of snapshot collections associated with volume collection.
	SnapcollCount int64 `json:"snapcoll_count,omitempty"`
	// ScheduleList List of snapshot schedules associated with volume collection.
	ScheduleList []any `json:"schedule_list,omitempty"`
	// ReplicationPartner Replication partner for this volume collection.
	ReplicationPartner string `json:"replication_partner,omitempty"`
	// LastReplicatedSnapcoll Last replicated snapshot collection on this volume collection.
	LastReplicatedSnapcoll string `json:"last_replicated_snapcoll,omitempty"`
	// LastReplicatedSnapcollList List of snapshot collection information for the last replicated snapshot collection per schedule.
	LastReplicatedSnapcollList []any `json:"last_replicated_snapcoll_list,omitempty"`
	// ProtectionType Specifies if volume collection is protected with schedules.
	ProtectionType bool `json:"protection_type,omitempty"`
	// LagTime Replication lag time for volume collection.
	LagTime int64 `json:"lag_time,omitempty"`
	// IsStandaloneVolcoll Indicates whether this is a standalone volume collection.
	IsStandaloneVolcoll bool `json:"is_standalone_volcoll,omitempty"`
	// TotalReplBytes Total size of volumes to be replicated for this volume collection.
	TotalReplBytes int64 `json:"total_repl_bytes,omitempty"`
	// ReplBytesTransferred Total size of volumes replicated for this volume collection.
	ReplBytesTransferred int64 `json:"repl_bytes_transferred,omitempty"`
	// IsHandingOver Indicates whether a handover operation is in progress on this volume collection.
	IsHandingOver bool `json:"is_handing_over,omitempty"`
	// HandoverReplicationPartner Replication partner to which ownership is being transferred as part of handover operation.
	HandoverReplicationPartner string `json:"handover_replication_partner,omitempty"`
	// Metadata Key-value pairs that augment a volume collection's attributes.
	Metadata []any `json:"metadata,omitempty"`
	// SrepLastSync Time when a synchronously replicated volume collection was last synchronized.
	SrepLastSync string `json:"srep_last_sync,omitempty"`
	// SrepResyncPercent Percentage of the resync progress for a synchronously replicated volume collection.
	SrepResyncPercent int64 `json:"srep_resync_percent,omitempty"`
}

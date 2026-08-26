package models

// SnapshotCollection models the NimbleOS snapshot_collections object set.
type SnapshotCollection struct {
	// ID Identifier for snapshot collection.
	ID string `json:"id,omitempty"`
	// Name Name of snapshot collection.
	Name string `json:"name,omitempty"`
	// Description Text description of snapshot collection.
	Description string `json:"description,omitempty"`
	// VolcollName Volume collection name.
	VolcollName string `json:"volcoll_name,omitempty"`
	// VolcollID Parent volume collection ID.
	VolcollID string `json:"volcoll_id,omitempty"`
	// OriginName Origination group name/ID.
	OriginName string `json:"origin_name,omitempty"`
	// IsReplica Indicates if snapshot collection was created as a replica.
	IsReplica bool `json:"is_replica,omitempty"`
	// SrepOwnerName Name of the partner where the snapshots in this snapshot collection reside.
	SrepOwnerName string `json:"srep_owner_name,omitempty"`
	// SrepOwnerID ID of the partner where snapshots for this snapshot collection reside which were created by synchronous replication.
	SrepOwnerID string `json:"srep_owner_id,omitempty"`
	// PeerSnapcollID ID of the peer snapshot collection created by synchronous replication.
	PeerSnapcollID string `json:"peer_snapcoll_id,omitempty"`
	// NumSnaps Current number of live, non-hidden snaps in this collection.
	NumSnaps string `json:"num_snaps,omitempty"`
	// IsComplete Is complete.
	IsComplete bool `json:"is_complete,omitempty"`
	// IsManual Is manual.
	IsManual bool `json:"is_manual,omitempty"`
	// IsExternalTrigger Is externally triggered.
	IsExternalTrigger bool `json:"is_external_trigger,omitempty"`
	// IsUnmanaged Indicates whether a snapshot collection is unmanaged.
	IsUnmanaged bool `json:"is_unmanaged,omitempty"`
	// IsManuallyManaged Indicates whether a snapshot collection is managed.
	IsManuallyManaged bool `json:"is_manually_managed,omitempty"`
	// ReplStatus Replication status.
	ReplStatus string `json:"repl_status,omitempty"`
	// ReplStartTime Replication start time.
	ReplStartTime int64 `json:"repl_start_time,omitempty"`
	// ReplCompleteTime Replication complete time.
	ReplCompleteTime int64 `json:"repl_complete_time,omitempty"`
	// ReplBytesTransferred Number of bytes transferred after replication completes.
	ReplBytesTransferred int64 `json:"repl_bytes_transferred,omitempty"`
	// CreationTime Time when this snapshot collection was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this snapshot collection was last modified.
	LastModified string `json:"last_modified,omitempty"`
	// OnlineStatus Online status of snapcoll.
	OnlineStatus bool `json:"online_status,omitempty"`
	// VolSnapAttrList List of snapshot attributes for snapshots being created as part of snapshot collection creation.
	VolSnapAttrList []any `json:"vol_snap_attr_list,omitempty"`
	// SnapshotsList List of snapshots in the snapshot collection.
	SnapshotsList []any `json:"snapshots_list,omitempty"`
	// Replicate True if this snapshot collection has been marked for replication.
	Replicate bool `json:"replicate,omitempty"`
	// ReplicateTo Specifies the partner name that the snapshots in this snapshot collection are replicated to.
	ReplicateTo bool `json:"replicate_to,omitempty"`
	// StartOnline Start with snapshot set online.
	StartOnline string `json:"start_online,omitempty"`
	// AllowWrites Allow applications to write to created snapshot(s).
	AllowWrites bool `json:"allow_writes,omitempty"`
	// DisableAppsync Do not perform application synchronization for this snapshot, create a crash-consistent snapshot instead.
	DisableAppsync bool `json:"disable_appsync,omitempty"`
	// SnapVerify Run verification tool on this snapshot.
	SnapVerify string `json:"snap_verify,omitempty"`
	// SkipDbConsistencyCheck Skip consistency check for database files on this snapshot.
	SkipDbConsistencyCheck bool `json:"skip_db_consistency_check,omitempty"`
	// SchedID ID of protection schedule of snapshot collection.
	SchedID string `json:"sched_id,omitempty"`
	// SchedName Name of protection schedule of snapshot collection.
	SchedName string `json:"sched_name,omitempty"`
	// InvokeOnUpstreamPartner Invoke snapshot request on upstream partner.
	InvokeOnUpstreamPartner bool `json:"invoke_on_upstream_partner,omitempty"`
	// AgentType External management agent type for snapshots being created as part of snapshot collection.
	AgentType int64 `json:"agent_type,omitempty"`
	// ExpiryAfter Number of seconds after which this snapcoll is considered expired by the snapshot TTL.
	ExpiryAfter string `json:"expiry_after,omitempty"`
	// Metadata Key-value pairs that augment a snapshot collection's attributes.
	Metadata []any `json:"metadata,omitempty"`
	// Force Forcibly delete the specified snapshot collection even if it is the last replicated snapshot.
	Force bool `json:"force,omitempty"`
}

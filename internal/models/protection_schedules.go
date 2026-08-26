package models

// ProtectionSchedule models the NimbleOS protection_schedules object set.
type ProtectionSchedule struct {
	// ID Identifier for protection schedule.
	ID string `json:"id,omitempty"`
	// Name Name of snapshot schedule to create.
	Name string `json:"name,omitempty"`
	// Description Description of the schedule.
	Description string `json:"description,omitempty"`
	// VolcollOrProttmplType Type of the protection policy this schedule is attached to.
	VolcollOrProttmplType string `json:"volcoll_or_prottmpl_type,omitempty"`
	// VolcollOrProttmplID Identifier of the protection policy (protection_template or volume_collection) in which this protection schedule is attached to.
	VolcollOrProttmplID string `json:"volcoll_or_prottmpl_id,omitempty"`
	// Period Repeat interval for snapshots with respect to the period_unit.
	Period int64 `json:"period,omitempty"`
	// PeriodUnit Time unit over which to take the number of snapshots specified in 'period'.
	PeriodUnit int64 `json:"period_unit,omitempty"`
	// AtTime Time of day when snapshot should be taken.
	AtTime int64 `json:"at_time,omitempty"`
	// UntilTime Time of day to stop taking snapshots.
	UntilTime int64 `json:"until_time,omitempty"`
	// Days Specifies which days snapshots should be taken.
	Days []string `json:"days,omitempty"`
	// NumRetain Number of snapshots to retain.
	NumRetain string `json:"num_retain,omitempty"`
	// DownstreamPartner Specifies the partner name if snapshots created by this schedule should be replicated.
	DownstreamPartner string `json:"downstream_partner,omitempty"`
	// DownstreamPartnerName Specifies the partner name if snapshots created by this schedule should be replicated.
	DownstreamPartnerName string `json:"downstream_partner_name,omitempty"`
	// DownstreamPartnerID Specifies the partner ID if snapshots created by this schedule should be replicated.
	DownstreamPartnerID string `json:"downstream_partner_id,omitempty"`
	// UpstreamPartnerName Specifies the partner name from which snapshots created by this schedule are replicated.
	UpstreamPartnerName string `json:"upstream_partner_name,omitempty"`
	// UpstreamPartnerID Specifies the partner ID from which snapshots created by this schedule are replicated.
	UpstreamPartnerID string `json:"upstream_partner_id,omitempty"`
	// ReplicateEvery Specifies which snapshots should be replicated.
	ReplicateEvery bool `json:"replicate_every,omitempty"`
	// NumRetainReplica Number of snapshots to retain on the replica.
	NumRetainReplica string `json:"num_retain_replica,omitempty"`
	// ReplAlertThres Replication alert threshold in seconds.
	ReplAlertThres string `json:"repl_alert_thres,omitempty"`
	// SnapVerify Run verification tool on snapshot created by this schedule.
	SnapVerify string `json:"snap_verify,omitempty"`
	// SkipDbConsistencyCheck Skip consistency check for database files on snapshots created by this schedule.
	SkipDbConsistencyCheck bool `json:"skip_db_consistency_check,omitempty"`
	// DisableAppsync Disables application synchronized snapshots and creates crash consistent snapshots instead.
	DisableAppsync bool `json:"disable_appsync,omitempty"`
	// ScheduleType Normal schedules have internal timers which drive snapshot creation.
	ScheduleType string `json:"schedule_type,omitempty"`
	// Active A schedule is active only if it is owned by the same owner as the volume collection.
	Active bool `json:"active,omitempty"`
	// CreationTime Time when this protection schedule was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this protection schedule was last modified.
	LastModified string `json:"last_modified,omitempty"`
	// LastModSchedTime Time when the timing of the protection schedule was last modified.
	LastModSchedTime int64 `json:"last_mod_sched_time,omitempty"`
	// LastReplicatedSnapcollName Specifies the name of last replicated snapshot collection.
	LastReplicatedSnapcollName string `json:"last_replicated_snapcoll_name,omitempty"`
	// LastReplicatedSnapcollID Specifies the snapshot collection ID of last replicated snapshot collection.
	LastReplicatedSnapcollID string `json:"last_replicated_snapcoll_id,omitempty"`
	// LastReplicatedAtTime Time when last snapshot collection was replicated.
	LastReplicatedAtTime int64 `json:"last_replicated_at_time,omitempty"`
	// LastSnapTime Time when last snapshot was taken.
	LastSnapTime int64 `json:"last_snap_time,omitempty"`
	// NextSnapTime Time when next snapshot will be taken.
	NextSnapTime int64 `json:"next_snap_time,omitempty"`
	// NextReplSnapTime Time when next snapshot will be replicated.
	NextReplSnapTime int64 `json:"next_repl_snap_time,omitempty"`
	// SnapCounter This is only used by custom read handler for internal calculations.
	SnapCounter int64 `json:"snap_counter,omitempty"`
	// SchedOwnerID Identifier of the group that owns this schedule.
	SchedOwnerID string `json:"sched_owner_id,omitempty"`
	// SchedOwnerName Name of the group that owns this schedule.
	SchedOwnerName string `json:"sched_owner_name,omitempty"`
	// LastConfigChangeTime The last timing configutation changed.
	LastConfigChangeTime int64 `json:"last_config_change_time,omitempty"`
	// CurrentlyReplicatingSnapcollName The name of the currently replicating snapshot collection if one exists, the empty string otherwise.
	CurrentlyReplicatingSnapcollName string `json:"currently_replicating_snapcoll_name,omitempty"`
	// VolStatusList The list of the replication status of volumes undergoing replication.
	VolStatusList []any `json:"vol_status_list,omitempty"`
	// SyncReplVolStatusList A list of the replication status of volumes undergoing synchronous replication.
	SyncReplVolStatusList []any `json:"sync_repl_vol_status_list,omitempty"`
	// UseDownstreamForDR Break synchronous replication for the specified volume collection and present downstream volumes to host(s).
	UseDownstreamForDR string `json:"use_downstream_for_DR,omitempty"`
}

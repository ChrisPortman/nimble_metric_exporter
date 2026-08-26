package models

// ProtectionTemplate models the NimbleOS protection_templates object set.
type ProtectionTemplate struct {
	// ID Identifier for protection template.
	ID string `json:"id,omitempty"`
	// Name User provided identifier.
	Name string `json:"name,omitempty"`
	// FullName Fully qualified name of protection template.
	FullName string `json:"full_name,omitempty"`
	// SearchName Name of protection template used for object search.
	SearchName string `json:"search_name,omitempty"`
	// Description Text description of protection template.
	Description string `json:"description,omitempty"`
	// ReplPriority Replication priority for the protection template with the following choices: {normal | high}.
	ReplPriority int64 `json:"repl_priority,omitempty"`
	// AppSync Application synchronization ({none|vss|vmware|generic}).
	AppSync string `json:"app_sync,omitempty"`
	// AppServer Application server hostname.
	AppServer string `json:"app_server,omitempty"`
	// AppID Application ID running on the server.
	AppID string `json:"app_id,omitempty"`
	// AppClusterName If the application is running within a Windows cluster environment then this is the cluster name.
	AppClusterName string `json:"app_cluster_name,omitempty"`
	// AppServiceName If the application is running within a Windows cluster environment then this is the instance name of the service running within the cluster environment.
	AppServiceName string `json:"app_service_name,omitempty"`
	// VcenterHostname VMware vCenter hostname.
	VcenterHostname string `json:"vcenter_hostname,omitempty"`
	// VcenterUsername VMware vCenter username.
	VcenterUsername string `json:"vcenter_username,omitempty"`
	// VcenterPassword VMware vCenter password.
	VcenterPassword string `json:"vcenter_password,omitempty"`
	// AgentHostname Generic Backup agent hostname.
	AgentHostname int64 `json:"agent_hostname,omitempty"`
	// AgentUsername Generic Backup agent username.
	AgentUsername int64 `json:"agent_username,omitempty"`
	// AgentPassword Generic Backup agent password.
	AgentPassword int64 `json:"agent_password,omitempty"`
	// CreationTime Time when this protection template was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this protection template was last modified.
	LastModified string `json:"last_modified,omitempty"`
	// ScheduleList List of schedules for this protection policy.
	ScheduleList []any `json:"schedule_list,omitempty"`
}

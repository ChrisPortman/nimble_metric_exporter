package models

// NetworkConfig models the NimbleOS network_configs object set.
type NetworkConfig struct {
	// ID Identifier for network configuration.
	ID string `json:"id,omitempty"`
	// Name Name of the network configuration.
	Name string `json:"name,omitempty"`
	// MgmtIP Management IP address for the Group.
	MgmtIP string `json:"mgmt_ip,omitempty"`
	// SecondaryMgmtIP Secondary management IP address for the Group.
	SecondaryMgmtIP string `json:"secondary_mgmt_ip,omitempty"`
	// Role Role of network configuration.
	Role string `json:"role,omitempty"`
	// ISCSIAutomaticConnectionMethod Whether automatic connection method is enabled.
	ISCSIAutomaticConnectionMethod bool `json:"iscsi_automatic_connection_method,omitempty"`
	// ISCSIConnectionRebalancing Whether rebalancing is enabled.
	ISCSIConnectionRebalancing bool `json:"iscsi_connection_rebalancing,omitempty"`
	// RouteList List of static routes.
	RouteList []any `json:"route_list,omitempty"`
	// SubnetList List of subnet configs.
	SubnetList []any `json:"subnet_list,omitempty"`
	// ArrayList List of array network configs.
	ArrayList []any `json:"array_list,omitempty"`
	// GroupLeaderArray Name of the group leader array.
	GroupLeaderArray string `json:"group_leader_array,omitempty"`
	// CreationTime Time when this net configuration was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this network configuration was last modified.
	LastModified string `json:"last_modified,omitempty"`
	// ActiveSince Start time of activity.
	ActiveSince bool `json:"active_since,omitempty"`
	// LastActive Time of last activity.
	LastActive string `json:"last_active,omitempty"`
	// IgnoreValidationMask Indicates whether to ignore the validation.
	IgnoreValidationMask bool `json:"ignore_validation_mask,omitempty"`
}

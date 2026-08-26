package models

// Subnet models the NimbleOS subnets object set.
type Subnet struct {
	// ID Identifier for the initiator group.
	ID string `json:"id,omitempty"`
	// Name Name of subnet configuration.
	Name string `json:"name,omitempty"`
	// Network Subnet network address.
	Network string `json:"network,omitempty"`
	// Netmask Subnet netmask address.
	Netmask string `json:"netmask,omitempty"`
	// Type Subnet type.
	Type string `json:"type,omitempty"`
	// AllowISCSI Subnet type.
	AllowISCSI bool `json:"allow_iscsi,omitempty"`
	// AllowGroup Subnet type.
	AllowGroup bool `json:"allow_group,omitempty"`
	// DiscoveryIP Subnet network address.
	DiscoveryIP string `json:"discovery_ip,omitempty"`
	// MTU MTU for specified subnet.
	MTU string `json:"mtu,omitempty"`
	// NetzoneType Specify Network Affinity Zone type for iSCSI enabled subnets.
	NetzoneType string `json:"netzone_type,omitempty"`
	// VlanID VLAN ID for specified subnet.
	VlanID string `json:"vlan_id,omitempty"`
	// CreationTime Time when this subnet configuration was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this subnet configuration was last modified.
	LastModified string `json:"last_modified,omitempty"`
	// Failover Failover setting of the subnet.
	Failover string `json:"failover,omitempty"`
	// FailoverEnableTime Failover for this subnet will be enabled again at the time specified by failover_enable_time.
	FailoverEnableTime int64 `json:"failover_enable_time,omitempty"`
}

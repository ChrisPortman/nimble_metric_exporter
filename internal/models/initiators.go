package models

// Initiator models the NimbleOS initiators object set.
type Initiator struct {
	// ID Identifier for initiator.
	ID string `json:"id,omitempty"`
	// AccessProtocol Access protocol used by the initiator.
	AccessProtocol string `json:"access_protocol,omitempty"`
	// InitiatorGroupID Identifier of the initiator group that this initiator is assigned to.
	InitiatorGroupID string `json:"initiator_group_id,omitempty"`
	// InitiatorGroupName Name of the initiator group that this initiator is assigned to.
	InitiatorGroupName string `json:"initiator_group_name,omitempty"`
	// Label Unique Identifier of the iSCSI initiator.
	Label string `json:"label,omitempty"`
	// IQN IQN name of the iSCSI initiator.
	IQN string `json:"iqn,omitempty"`
	// IPAddress IP address of the iSCSI initiator.
	IPAddress string `json:"ip_address,omitempty"`
	// Alias Alias of the Fibre Channel initiator.
	Alias string `json:"alias,omitempty"`
	// ChapuserID Identifier for the CHAP user.
	ChapuserID string `json:"chapuser_id,omitempty"`
	// WWPN WWPN (World Wide Port Name) of the Fibre Channel initiator.
	WWPN string `json:"wwpn,omitempty"`
	// VpOverride Flag to allow modifying VP created initiator groups.
	VpOverride bool `json:"vp_override,omitempty"`
	// CreationTime Time when this initiator group was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this initiator group was last modified.
	LastModified string `json:"last_modified,omitempty"`
	// OverrideExistingAlias Forcibly add Fibre Channel initiator to initiator group by updating or removing conflicting Fibre Channel initiator aliases.
	OverrideExistingAlias bool `json:"override_existing_alias,omitempty"`
}

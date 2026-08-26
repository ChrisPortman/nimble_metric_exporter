package models

// FibreChannelConfig models the NimbleOS fibre_channel_configs object set.
type FibreChannelConfig struct {
	// ID Identifier for Fibre Channel configuration.
	ID string `json:"id,omitempty"`
	// ArrayList List of array Fibre Channel configs.
	ArrayList []any `json:"array_list,omitempty"`
	// GroupLeaderArray Name of the group leader array.
	GroupLeaderArray string `json:"group_leader_array,omitempty"`
}

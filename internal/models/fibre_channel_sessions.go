package models

// FibreChannelSession models the NimbleOS fibre_channel_sessions object set.
type FibreChannelSession struct {
	// ID Unique identifier of the Fibre Channel session.
	ID string `json:"id,omitempty"`
	// InitiatorInfo Information about the Fibre Channel initiator.
	InitiatorInfo string `json:"initiator_info,omitempty"`
	// TargetInfo Information about the Fibre Channel target.
	TargetInfo string `json:"target_info,omitempty"`
}

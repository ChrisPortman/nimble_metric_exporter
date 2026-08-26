package models

// FibreChannelInitiatorAlias models the NimbleOS fibre_channel_initiator_aliases object set.
type FibreChannelInitiatorAlias struct {
	// ID Unique identifier for the Fibre Channel initiator alias.
	ID string `json:"id,omitempty"`
	// Alias Alias of the Fibre Channel initiator.
	Alias string `json:"alias,omitempty"`
	// WWPN WWPN (World Wide Port Name) of the Fibre Channel initiator.
	WWPN string `json:"wwpn,omitempty"`
	// Source Source of the Fibre Channel initiator alias.
	Source string `json:"source,omitempty"`
}

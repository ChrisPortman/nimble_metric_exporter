package models

// NetworkInterface models the NimbleOS network_interfaces object set.
type NetworkInterface struct {
	// ID Identifier for the interface.
	ID string `json:"id,omitempty"`
	// ArrayNameOrSerial Name or serial of the array where the interface is hosted.
	ArrayNameOrSerial string `json:"array_name_or_serial,omitempty"`
	// PartialResponseOk Indicate that it is ok to provide partially available response.
	PartialResponseOk string `json:"partial_response_ok,omitempty"`
	// ArrayID Identifier for the array.
	ArrayID string `json:"array_id,omitempty"`
	// ControllerName Name (A or B) of the controller where the interface is hosted.
	ControllerName string `json:"controller_name,omitempty"`
	// ControllerID Identifier of the controller where the interface is hosted.
	ControllerID string `json:"controller_id,omitempty"`
	// Name Name of the interface.
	Name string `json:"name,omitempty"`
	// MAC MAC address of the interface.
	MAC string `json:"mac,omitempty"`
	// IsPresent Whether this interface is present on this controller.
	IsPresent bool `json:"is_present,omitempty"`
	// LinkSpeed Speed of the link.
	LinkSpeed int64 `json:"link_speed,omitempty"`
	// LinkStatus Status of the link.
	LinkStatus string `json:"link_status,omitempty"`
	// MTU MTU on the link.
	MTU string `json:"mtu,omitempty"`
	// Port Port number for this interface.
	Port int64 `json:"port,omitempty"`
	// Slot Slot number for this interface.
	Slot int64 `json:"slot,omitempty"`
	// MaxLinkSpeed Maximum speed of the link.
	MaxLinkSpeed int64 `json:"max_link_speed,omitempty"`
	// NICType Interface type.
	NICType string `json:"nic_type,omitempty"`
	// IPList List of IP addresses assigned to this network interface.
	IPList []any `json:"ip_list,omitempty"`
}

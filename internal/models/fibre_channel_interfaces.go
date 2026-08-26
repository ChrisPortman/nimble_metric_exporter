package models

// FibreChannelInterface models the NimbleOS fibre_channel_interfaces object set.
type FibreChannelInterface struct {
	// ID Identifier for the Fibre Channel interface.
	ID string `json:"id,omitempty"`
	// ArrayNameOrSerial Name or serial number of array where the interface is hosted.
	ArrayNameOrSerial string `json:"array_name_or_serial,omitempty"`
	// PartialResponseOk Indicate that it is ok to provide partially available response.
	PartialResponseOk string `json:"partial_response_ok,omitempty"`
	// ControllerName Name (A or B) of the controller where the interface is hosted.
	ControllerName string `json:"controller_name,omitempty"`
	// FCPortID ID of the port with which the interface is associated.
	FCPortID int64 `json:"fc_port_id,omitempty"`
	// Name Name of Fibre Channel interface.
	Name string `json:"name,omitempty"`
	// WWNN WWNN (World Wide Node Name) for this Fibre Channel interface.
	WWNN string `json:"wwnn,omitempty"`
	// WWPN WWPN (World Wide Port Name) for this Fibre Channel interface.
	WWPN string `json:"wwpn,omitempty"`
	// Peerzone Active peer zone for this Fibre Channel interface.
	Peerzone string `json:"peerzone,omitempty"`
	// Online Identify whether the Fibre Channel interface is online.
	Online bool `json:"online,omitempty"`
	// FirmwareVersion Version of the Fibre Channel firmware.
	FirmwareVersion string `json:"firmware_version,omitempty"`
	// LogicalPortNumber Logical port number for the Fibre Channel port.
	LogicalPortNumber int64 `json:"logical_port_number,omitempty"`
	// FCPortName Name of Fibre Channel port.
	FCPortName int64 `json:"fc_port_name,omitempty"`
	// BusLocation PCI bus location of the HBA for this Fibre Channel port.
	BusLocation string `json:"bus_location,omitempty"`
	// Slot HBA slot number for this Fibre Channel port.
	Slot int64 `json:"slot,omitempty"`
	// Orientation Orientation of FC ports on a HBA.
	Orientation string `json:"orientation,omitempty"`
	// Port HBA port number for this Fibre Channel port.
	Port int64 `json:"port,omitempty"`
	// LinkInfo Information about the Fibre Channel link at which interface is operating.
	LinkInfo map[string]any `json:"link_info,omitempty"`
	// FabricInfo Fibre Channel fabric information.
	FabricInfo map[string]any `json:"fabric_info,omitempty"`
}

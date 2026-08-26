package models

// FibreChannelPort models the NimbleOS fibre_channel_ports object set.
type FibreChannelPort struct {
	// ID Identifier for the Fibre Channel port.
	ID string `json:"id,omitempty"`
	// ArrayNameOrSerial Name or serial number of the array.
	ArrayNameOrSerial string `json:"array_name_or_serial,omitempty"`
	// ControllerName Name (A or B) of the controller to which the port belongs.
	ControllerName string `json:"controller_name,omitempty"`
	// FCPortName Name of the Fibre Channel port.
	FCPortName int64 `json:"fc_port_name,omitempty"`
	// BusLocation PCI bus location of the HBA (Host Bus Adapter) for this Fibre Channel port.
	BusLocation string `json:"bus_location,omitempty"`
	// Port HBA (Host Bus Adapter) port number for this Fibre Channel port.
	Port int64 `json:"port,omitempty"`
	// Slot HBA (Host Bus Adapter) slot number for this Fibre Channel port.
	Slot int64 `json:"slot,omitempty"`
	// Orientation Orientation of FC ports on a HBA.
	Orientation string `json:"orientation,omitempty"`
	// LinkInfo Information about the Fibre Channel link associated with this port.
	LinkInfo map[string]any `json:"link_info,omitempty"`
	// RxPower SFP RX power in uW.
	RxPower string `json:"rx_power,omitempty"`
	// TxPower SFP TX power in uW.
	TxPower string `json:"tx_power,omitempty"`
}

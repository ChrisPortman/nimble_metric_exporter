package models

// Controller models the NimbleOS controllers object set.
type Controller struct {
	// ID Identifier of the controller.
	ID string `json:"id,omitempty"`
	// Name Name of the controller.
	Name string `json:"name,omitempty"`
	// ArrayName Name of the array containing this controller.
	ArrayName string `json:"array_name,omitempty"`
	// ArrayID Rest ID of the array containing this controller.
	ArrayID string `json:"array_id,omitempty"`
	// PartialResponseOk Indicate that it is ok to provide partially available response.
	PartialResponseOk string `json:"partial_response_ok,omitempty"`
	// Serial Serial number for this controller.
	Serial string `json:"serial,omitempty"`
	// Model Model of this controller.
	Model string `json:"model,omitempty"`
	// Hostname Host name for the controller.
	Hostname string `json:"hostname,omitempty"`
	// SupportAddress IP address used for support.
	SupportAddress int64 `json:"support_address,omitempty"`
	// SupportNetmask IP netmask used for support.
	SupportNetmask int64 `json:"support_netmask,omitempty"`
	// SupportNIC Network card used for support.
	SupportNIC int64 `json:"support_nic,omitempty"`
	// PowerStatus Overall power supply status for the controller.
	PowerStatus string `json:"power_status,omitempty"`
	// FanStatus Overall fan status for the controller.
	FanStatus string `json:"fan_status,omitempty"`
	// TemperatureStatus Overall temperature status for the controller.
	TemperatureStatus string `json:"temperature_status,omitempty"`
	// PowerSupplies Status for each power supply in the controller.
	PowerSupplies []any `json:"power_supplies,omitempty"`
	// Fans Status for each fan in the controller.
	Fans []any `json:"fans,omitempty"`
	// TemperatureSensors Status for temperature sensor in the controller.
	TemperatureSensors []any `json:"temperature_sensors,omitempty"`
	// PartitionStatus Status of the system's raid partitions.
	PartitionStatus string `json:"partition_status,omitempty"`
	// CtrlrSide Identifies which controller this is on its array.
	CtrlrSide string `json:"ctrlr_side,omitempty"`
	// State Indicates whether this controller is active or not.
	State bool `json:"state,omitempty"`
	// NVMeCardsEnabled Indicates if the NVMe accelerator card is enabled.
	NVMeCardsEnabled string `json:"nvme_cards_enabled,omitempty"`
	// NVMeCards List of NVMe accelerator cards.
	NVMeCards []any `json:"nvme_cards,omitempty"`
	// AsupTime Time of the last autosupport by the controller.
	AsupTime int64 `json:"asup_time,omitempty"`
}

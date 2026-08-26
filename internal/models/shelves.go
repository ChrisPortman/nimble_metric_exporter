package models

// Shelf models the NimbleOS shelves object set.
type Shelf struct {
	// ID ID of shelf.
	ID string `json:"id,omitempty"`
	// ArrayName Name of array the shelf belongs to.
	ArrayName string `json:"array_name,omitempty"`
	// ArrayID ID of array the shelf belongs to.
	ArrayID string `json:"array_id,omitempty"`
	// PartialResponseOk Indicate that it is okay to provide partially available response.
	PartialResponseOk string `json:"partial_response_ok,omitempty"`
	// ChassisType Chassis type. One of 'chassis_unknown', 'chassis_3u16', 'chassis_4u24'
	ChassisType string `json:"chassis_type,omitempty"`
	// Ctrlrs List of ctrlr info.
	Ctrlrs []ShelfController `json:"ctrlrs,omitempty"`
	// Serial The serial number of the chassis.
	Serial string `json:"serial,omitempty"`
	// Model Model of the shelf or head unit.
	Model string `json:"model,omitempty"`
	// ModelExt Extended model of the shelf or head unit.
	ModelExt string `json:"model_ext,omitempty"`
	// ChassisSensors List of chassis sensor readings.
	ChassisSensors []ShelfSensor `json:"chassis_sensors,omitempty"`
	// PSUOverallStatus The overall status for the PSUs. One of 'OK', 'Alerted', 'Failed', 'Missing'.
	PSUOverallStatus string `json:"psu_overall_status,omitempty"`
	// FanOverallStatus The overall status for the fans on both controllers. One of 'OK', 'Alerted', 'Failed', 'Missing'.
	FanOverallStatus string `json:"fan_overall_status,omitempty"`
	// TempOverallStatus The overall status for the temperature on both controllers. One of 'OK', 'Alerted', 'Failed', 'Missing'.
	TempOverallStatus string `json:"temp_overall_status,omitempty"`
	// DiskSets Attributes for the disk sets in this shelf.
	DiskSets []any `json:"disk_sets,omitempty"`
	// Activated Activated state for shelf or disk set means it is available to store date on.
	Activated bool `json:"activated,omitempty"`
	// Driveset Driveset to activate.
	Driveset string `json:"driveset,omitempty"`
	// Force Forcibly activate shelf.
	Force bool `json:"force,omitempty"`
	// AcceptForeign Accept the removal of data on the shelf disks and activate foreign shelf.
	AcceptForeign bool `json:"accept_foreign,omitempty"`
	// AcceptDedupeImpact Accept the reduction or elimination of deduplication capability on the system as a result of activating a shelf that does not meet the necessary deduplication requirements.
	AcceptDedupeImpact bool `json:"accept_dedupe_impact,omitempty"`
	// LastRequest Indicates this is the last request in a series of shelf add requests.
	LastRequest string `json:"last_request,omitempty"`
}

// ShelfController model is returned within shelf data structures
type ShelfController struct {
	CachedSerial          string        `json:"cached_serial,omitempty"`
	ControllerAttrSetList []any         `json:"ctrlr_attrset_list,omitempty"`
	HardwareModel         string        `json:"ctrlr_hw_model,omitempty"`
	SensorLastRun         int64         `json:"ctrlr_sensor_last_run,omitempty"`
	Sensors               []ShelfSensor `json:"ctrlr_sensors,omitempty"`
	Side                  string        `json:"ctrlr_side,omitempty"`
	EncLocId              int64         `json:"enc_loc_id,omitempty"`
	ExpandedSASAddr       string        `json:"exp_sas_addr,omitempty"`
	FanOverallStatus      string        `json:"fan_overall_status,omitempty"`
	HwMasterState         string        `json:"hw_master_state,omitempty"`
	HwMembershipFailure   bool          `json:"hw_mship_failure,omitempty"`
	PortInfo              []any         `json:"port_info,omitempty"`
	SwMasterState         string        `json:"sw_master_state,omitempty"`
	TempOverallStatus     string        `json:"temp_overall_status,omitempty"`
}

// ShelfSensor model describe the sensors that appear in shelf responses
type ShelfSensor struct {
	// CID identifies the chassis
	CID string `json:"cid,omitempty"`
	// DisplayName of the sensor
	DisplayName string `json:"display_name,omitempty"`
	// Location of the sensor
	Location string `json:"Location,omitempty"`
	// Name of the sensor
	Name string `json:"name,omitempty"`
	// Status of the sensor. "OK" or something else.
	Status string `json:"status,omitempty"`
	// Type of sensor
	Type string `json:"type,omitempty"`
	// Value of sensor
	Value int64 `json:"value,omitempty"`
}

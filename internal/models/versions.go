package models

// Version models the NimbleOS versions object set.
type Version struct {
	// Name API version number.
	Name string `json:"name,omitempty"`
	// SoftwareVersion Software version of the group.
	SoftwareVersion string `json:"software_version,omitempty"`
}

package models

// SoftwareVersion models the NimbleOS software_versions object set.
type SoftwareVersion struct {
	// Version Software version, used as identifier in URL.
	Version string `json:"version,omitempty"`
	// Signature Keyed hash of download package.
	Signature string `json:"signature,omitempty"`
	// Name Name of version.
	Name string `json:"name,omitempty"`
	// Status Status of version.
	Status string `json:"status,omitempty"`
	// TotalBytes Size of version.
	TotalBytes int64 `json:"total_bytes,omitempty"`
	// DownloadedBytes Number of bytes downloaded for the version.
	DownloadedBytes int64 `json:"downloaded_bytes,omitempty"`
	// BlacklistReason Reason for blacklisting the version.
	BlacklistReason string `json:"blacklist_reason,omitempty"`
	// ReleaseDate Date when software version was released.
	ReleaseDate string `json:"release_date,omitempty"`
	// IsManuallyDownloaded Whether or not the version was downloaded manually.
	IsManuallyDownloaded bool `json:"is_manually_downloaded,omitempty"`
	// ReleaseStatus Release status of software version.
	ReleaseStatus string `json:"release_status,omitempty"`
	// NoPartialResponse Indicate that it is not ok to provide partially available response.
	NoPartialResponse bool `json:"no_partial_response,omitempty"`
}

package models

// UserGroup models the NimbleOS user_groups object set.
type UserGroup struct {
	// ID Identifier for the user group.
	ID string `json:"id,omitempty"`
	// Name Name of the user group.
	Name string `json:"name,omitempty"`
	// Description Description of the user group.
	Description string `json:"description,omitempty"`
	// RoleID Identifier for the user group's role.
	RoleID string `json:"role_id,omitempty"`
	// Role Role of the user.
	Role string `json:"role,omitempty"`
	// InactivityTimeout The amount of time that the user session is inactive before timing out.
	InactivityTimeout int64 `json:"inactivity_timeout,omitempty"`
	// CreationTime Time when this user was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this user was last modified.
	LastModified string `json:"last_modified,omitempty"`
	// Disabled User is currently disabled.
	Disabled bool `json:"disabled,omitempty"`
	// ExternalID External ID of the user group.
	ExternalID string `json:"external_id,omitempty"`
	// DomainID Identifier of the domain this user group belongs to.
	DomainID string `json:"domain_id,omitempty"`
	// DomainName Role of the user.
	DomainName string `json:"domain_name,omitempty"`
}

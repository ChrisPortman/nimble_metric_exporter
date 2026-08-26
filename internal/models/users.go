package models

// User models the NimbleOS users object set.
type User struct {
	// ID Identifier for the user.
	ID string `json:"id,omitempty"`
	// Name Name of the user.
	Name string `json:"name,omitempty"`
	// SearchName Name of the user used for object search.
	SearchName string `json:"search_name,omitempty"`
	// Description Description of the user.
	Description string `json:"description,omitempty"`
	// RoleID Identifier for the user's role.
	RoleID string `json:"role_id,omitempty"`
	// Role Role of the user.
	Role string `json:"role,omitempty"`
	// Password User's login password.
	Password string `json:"password,omitempty"`
	// AuthPassword Authorization password for changing password.
	AuthPassword string `json:"auth_password,omitempty"`
	// OTPType Type of One Time Password authentication in use.
	OTPType string `json:"otp_type,omitempty"`
	// OTPReset When sent as true, this causes a reset of the One Time Password secret for the user.
	OTPReset string `json:"otp_reset,omitempty"`
	// InactivityTimeout The amount of time that the user session is inactive before timing out.
	InactivityTimeout int64 `json:"inactivity_timeout,omitempty"`
	// CreationTime Time when this user was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this user was last modified.
	LastModified string `json:"last_modified,omitempty"`
	// FullName Fully qualified name of the user.
	FullName string `json:"full_name,omitempty"`
	// EmailAddr Email address of the user.
	EmailAddr string `json:"email_addr,omitempty"`
	// TenantID Identifier for the tenant.
	TenantID string `json:"tenant_id,omitempty"`
	// TenantKey Tenant secret key for encrypting the password.
	TenantKey string `json:"tenant_key,omitempty"`
	// Disabled User is currently disabled.
	Disabled bool `json:"disabled,omitempty"`
	// AuthLock User was locked due to failed logins.
	AuthLock bool `json:"auth_lock,omitempty"`
	// LastLogin Last login time.
	LastLogin bool `json:"last_login,omitempty"`
	// LastLogout Last logout time.
	LastLogout string `json:"last_logout,omitempty"`
	// LoggedIn User is currently logged in.
	LoggedIn bool `json:"logged_in,omitempty"`
}

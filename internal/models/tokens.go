package models

// Token models the NimbleOS tokens object set.
type Token struct {
	// ID Object identifier for the session token.
	ID string `json:"id,omitempty"`
	// SessionToken Token used for authentication.
	SessionToken string `json:"session_token,omitempty"`
	// Username User name for the session.
	Username string `json:"username,omitempty"`
	// Password Password for the user.
	Password string `json:"password,omitempty"`
	// OTPCode One time password code.
	OTPCode string `json:"otp_code,omitempty"`
	// AppName Application name.
	AppName string `json:"app_name,omitempty"`
	// SdkName SDK name.
	SdkName string `json:"sdk_name,omitempty"`
	// SourceIP IP address from which the session originates.
	SourceIP string `json:"source_ip,omitempty"`
	// CreationTime Time when this token was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this token was last modified.
	LastModified int64 `json:"last_modified,omitempty"`
	// ExpiryTime Time when this token will expire.
	ExpiryTime int64 `json:"expiry_time,omitempty"`
	// ServerUUID Non mandatory 36 character uuid returned by the server.
	ServerUUID string `json:"server_uuid,omitempty"`
	// GrantType OAuth grant type, currently only support 'urn:ietf:params:oauth:grant-type:jwt-bearer'.
	GrantType string `json:"grant_type,omitempty"`
	// Assertion OAuth assertion, currently expecting a JWT token.
	Assertion string `json:"assertion,omitempty"`
}

// CreateTokenRequest contains credentials for creating a NimbleOS session token.
type CreateTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	OTPCode  string `json:"otp_code,omitempty"`
	AppName  string `json:"app_name,omitempty"`
	SdkName  string `json:"sdk_name,omitempty"`
}

// OAuthTokenRequest contains the JWT bearer parameters for OAuth token creation.
type OAuthTokenRequest struct {
	GrantType string `json:"grant_type"`
	Assertion string `json:"assertion"`
}

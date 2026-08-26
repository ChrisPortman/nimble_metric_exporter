package models

// AuditLog models the NimbleOS audit_log object set.
type AuditLog struct {
	// ID Identifier for the audit log record.
	ID string `json:"id,omitempty"`
	// Type Identifier for type of audit log record.
	Type string `json:"type,omitempty"`
	// ObjectID Identifier of object operated upon.
	ObjectID string `json:"object_id,omitempty"`
	// ObjectName Name of object operated upon.
	ObjectName string `json:"object_name,omitempty"`
	// ObjectType Type of the object being operated upon.
	ObjectType string `json:"object_type,omitempty"`
	// Scope Scope within which object exists, for example, name of the array for a NIC.
	Scope string `json:"scope,omitempty"`
	// Time Time when this operation was performed.
	Time int64 `json:"time,omitempty"`
	// Status Status of the operation -- success or failure.
	Status string `json:"status,omitempty"`
	// ErrorCode If the operation has failed, this indicates the error code corresponding to the failure.
	ErrorCode string `json:"error_code,omitempty"`
	// UserID Identifier of the user who performed the operation.
	UserID string `json:"user_id,omitempty"`
	// UserName Username of the user who performed the operation.
	UserName string `json:"user_name,omitempty"`
	// UserFullName Full name of the user who performed the operation.
	UserFullName string `json:"user_full_name,omitempty"`
	// SourceIP IP address from where the operation request originated.
	SourceIP string `json:"source_ip,omitempty"`
	// ExtUserID The user id of an external user.
	ExtUserID string `json:"ext_user_id,omitempty"`
	// ExtUserGroupID The group ID of an external user.
	ExtUserGroupID string `json:"ext_user_group_id,omitempty"`
	// ExtUserGroupName The group name of an external user.
	ExtUserGroupName string `json:"ext_user_group_name,omitempty"`
	// AppName Name of application from where the operation request was issued, for example, pam, VSS Agent, etc.
	AppName string `json:"app_name,omitempty"`
	// AccessType Name of access on how the operation request was issued, for example, GUI, CLI or API.
	AccessType string `json:"access_type,omitempty"`
	// Category Category of the audit log record.
	Category string `json:"category,omitempty"`
	// ActivityType Type of activity performed, for example, create, update or delete.
	ActivityType string `json:"activity_type,omitempty"`
	// Activity Description of activity performed and recorded in audit log.
	Activity string `json:"activity,omitempty"`
}

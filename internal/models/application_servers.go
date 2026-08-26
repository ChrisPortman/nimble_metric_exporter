package models

// ApplicationServer models the NimbleOS application_servers object set.
type ApplicationServer struct {
	// ID Identifier for the application server.
	ID string `json:"id,omitempty"`
	// Name Name for the application server.
	Name string `json:"name,omitempty"`
	// Hostname Application server hostname.
	Hostname string `json:"hostname,omitempty"`
	// Port Application server port number.
	Port int64 `json:"port,omitempty"`
	// Username Application server username.
	Username string `json:"username,omitempty"`
	// Description Text description of application server.
	Description string `json:"description,omitempty"`
	// Password Application server password.
	Password string `json:"password,omitempty"`
	// ServerType Application server type ({invalid|vss|vmware|cisco|stack_vision|container_node}).
	ServerType string `json:"server_type,omitempty"`
	// Metadata Key-value pairs that augment an application server's attributes.
	Metadata []any `json:"metadata,omitempty"`
	// CreationTime Time when this application server was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified Time when this application server was last modified.
	LastModified string `json:"last_modified,omitempty"`
}

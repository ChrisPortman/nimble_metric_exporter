package models

// ActiveDirectoryMembership models the NimbleOS active_directory_memberships object set.
type ActiveDirectoryMembership struct {
	// ID Identifier for the Active Directory Domain.
	ID string `json:"id,omitempty"`
	// Description Description for the Active Directory Domain.
	Description string `json:"description,omitempty"`
	// Name Identifier for the Active Directory domain.
	Name string `json:"name,omitempty"`
	// Netbios Netbios name for the Active Directory domain.
	Netbios string `json:"netbios,omitempty"`
	// ServerList List of IP addresses or names for the backup domain controller.
	ServerList []string `json:"server_list,omitempty"`
	// ComputerName The name of the computer account in the domain controller.
	ComputerName string `json:"computer_name,omitempty"`
	// OrganizationalUnit The location for the computer account.
	OrganizationalUnit string `json:"organizational_unit,omitempty"`
	// User Name of the Activer Directory user with Administrator's privilege.
	User string `json:"user,omitempty"`
	// Password Password for the Active Directory user.
	Password string `json:"password,omitempty"`
	// Enabled Active Directory authentication is currently enabled.
	Enabled bool `json:"enabled,omitempty"`
}

package models

// LdapDomain models the NimbleOS ldap_domains object set.
type LdapDomain struct {
	// ID Identifier for the LDAP Domain.
	ID string `json:"id,omitempty"`
	// DomainName Domain name.
	DomainName string `json:"domain_name,omitempty"`
	// DomainDescription Description of the domain.
	DomainDescription string `json:"domain_description,omitempty"`
	// DomainEnabled Indicates whether the LDAP domain is currently active or not.
	DomainEnabled bool `json:"domain_enabled,omitempty"`
	// ServerUriList A set of up to 3 LDAP URIs.
	ServerUriList []any `json:"server_uri_list,omitempty"`
	// BindUser Full Distinguished Name of LDAP admin user.
	BindUser string `json:"bind_user,omitempty"`
	// BindPassword Password for the Full Distinguished Name of LDAP admin user.
	BindPassword string `json:"bind_password,omitempty"`
	// BaseDn The Distinguished Name of the base object from which to start all searches.
	BaseDn string `json:"base_dn,omitempty"`
	// UserSearchFilter Limit the results returned based on specific search criteria.
	UserSearchFilter string `json:"user_search_filter,omitempty"`
	// UserSearchBaseList A set of upto 10 Relative Distinguished Names, relative to the Base DN, from which to search for User objects.
	UserSearchBaseList []string `json:"user_search_base_list,omitempty"`
	// GroupSearchFilter Limit the results returned based on specific search criteria.
	GroupSearchFilter string `json:"group_search_filter,omitempty"`
	// GroupSearchBaseList A set of upto 10 Relative Distinguished Names, relative to the Base DN, from which to search for Group objects.
	GroupSearchBaseList []string `json:"group_search_base_list,omitempty"`
	// SchemaType Enum values are OpenLDAP or AD.
	SchemaType string `json:"schema_type,omitempty"`
}

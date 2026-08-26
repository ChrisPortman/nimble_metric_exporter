package models

// TrustedOauthIssuer models the NimbleOS trusted_oauth_issuers object set.
type TrustedOauthIssuer struct {
	// ID Identifier for the trusted oauth issuer record.
	ID string `json:"id,omitempty"`
	// Name Issuer ID string.
	Name string `json:"name,omitempty"`
	// JwksUrl The URL from which the device will download the public key set for signature verification.
	JwksUrl string `json:"jwks_url,omitempty"`
	// KeySet List of public keys for verifying signatures.
	KeySet []any `json:"key_set,omitempty"`
}

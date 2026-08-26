package models

// UserPolicy models the NimbleOS user_policies object set.
type UserPolicy struct {
	// ID Identifier for the security policy.
	ID string `json:"id,omitempty"`
	// AllowedAttempts Number of authentication attempts allowed before the user account is locked.
	AllowedAttempts int64 `json:"allowed_attempts,omitempty"`
	// MinLength Minimum length for user passwords.
	MinLength int64 `json:"min_length,omitempty"`
	// Upper Number of uppercase characters required in user passwords.
	Upper int64 `json:"upper,omitempty"`
	// Lower Number of lowercase characters required in user passwords.
	Lower int64 `json:"lower,omitempty"`
	// Digit Number of numerical characters required in user passwords.
	Digit int64 `json:"digit,omitempty"`
	// Special Number of special characters required in user passwords.
	Special int64 `json:"special,omitempty"`
	// PreviousDiff Number of characters that must be different from the previous password.
	PreviousDiff int64 `json:"previous_diff,omitempty"`
	// NoReuse Number of times that a password must change before you can reuse a previous password.
	NoReuse bool `json:"no_reuse,omitempty"`
	// MaxSessions Maximum number of sessions allowed for a group.
	MaxSessions int64 `json:"max_sessions,omitempty"`
}

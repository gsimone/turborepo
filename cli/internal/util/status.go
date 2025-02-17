package util

import "fmt"

// CachingStatus represents the api server's perspective
// on whether remote caching should be allowed
type CachingStatus int

const (
	// CachingStatusDisabled indicates that the server will not accept or serve artifacts
	CachingStatusDisabled CachingStatus = iota
	// CachingStatusEnabled indicates that the server will accept and serve artifacts
	CachingStatusEnabled
	// CachingStatusOverLimit indicates that a usage limit has been hit and the
	// server will temporarily not accept or server artifacts
	CachingStatusOverLimit
)

// CachingStatusFromString parses a raw string to a caching status enum value
func CachingStatusFromString(raw string) (CachingStatus, error) {
	switch raw {
	case "disabled":
		return CachingStatusDisabled, nil
	case "enabled":
		return CachingStatusEnabled, nil
	case "over_limit":
		return CachingStatusOverLimit, nil
	default:
		return CachingStatusDisabled, fmt.Errorf("unknown caching status: %v", raw)
	}
}

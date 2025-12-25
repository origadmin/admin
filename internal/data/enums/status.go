package enums

// Status defines a general-purpose status for various entities.
type Status int8

const (
	StatusInvalid  Status = 0
	StatusEnabled  Status = 1
	StatusDisabled Status = 2

	StatusUnknown  = StatusInvalid
	StatusActive   = StatusEnabled
	StatusInactive = StatusDisabled
	StatusFrozen   = StatusDisabled
)

// String returns the string representation of the status.
func (s Status) String() string {
	switch s {
	case StatusEnabled:
		return "enabled"
	case StatusDisabled:
		return "disabled"
	default:
		return "unknown"
	}
}

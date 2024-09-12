package evaluate

// Status is the resulting status codes
type Status int

// Constants for all status codes
const (
	StatusOK       = Status(0)
	StatusWarning  = Status(1)
	StatusCritical = Status(2)
	StatusUnknown  = Status(3)
)

// String returns the lexical representation of a status code
func (s Status) String() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusWarning:
		return "WARNING"
	case StatusCritical:
		return "CRITICAL"
	case StatusUnknown:
		fallthrough
	default:
		return "UNKNOWN"
	}
}

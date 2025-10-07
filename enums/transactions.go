package enums

type TransactionStatus int

const (
	Success TransactionStatus = iota + 1
	Failed
)

// String returns the string representation of TransactionStatus
func (ts TransactionStatus) String() string {
	switch ts {
	case Success:
		return "Success"
	case Failed:
		return "Failed"
	default:
		return "unknown"
	}
}

// IsValid checks if the TransactionStatus value is valid
func (ts TransactionStatus) IsValid() bool {
	return ts == Success || ts == Failed
}

// FromString converts a string to TransactionStatus
func FromString(s string) (TransactionStatus, bool) {
	switch s {
	case "Success":
		return Success, true
	case "Failed":
		return Failed, true
	default:
		return 0, false
	}
}
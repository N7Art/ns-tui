package models

// Mode represents the current mode of the application
type Mode int

const (
	// InsertMode is for searching and typing
	InsertMode Mode = iota
	// NormalMode is for vim-style navigation
	NormalMode
	// DetailMode is for viewing package details
	DetailMode
)

// String returns the string representation of the mode
func (m Mode) String() string {
	switch m {
	case InsertMode:
		return "INSERT"
	case NormalMode:
		return "NORMAL"
	case DetailMode:
		return "DETAIL"
	default:
		return "UNKNOWN"
	}
}

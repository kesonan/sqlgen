package spec

const (
	_                 = iota
	ActionCreate      // ActionCreate represents a create action.
	ActionRead        // ActionRead represents a read action.
	ActionUpdate      // ActionUpdate represents an update action.
	ActionDelete      // ActionDelete represents a delete action.
	ActionTransaction // ActionTransaction represents a transaction action.
)

// Action represents an action.
type Action int

package data

// CreateTimer is implemented by struct who wants to track Create Time
type CreateTimer interface {
	// CreateTime sets the creation time
	CreateTime()
}

// ModifyTimer is implemented by struct who wants to track modify time
type ModifyTimer interface {
	// ModifyTime sets the modification time
	ModifyTime()
}

// IDSetter is implemented by struct who wants to set ID
type IDSetter interface {
	SetID(string)
}

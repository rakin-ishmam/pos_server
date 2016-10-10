package data

// Access represents use access level
type Access string

const (
	// AccessRead defines read access
	AccessRead Access = "Read"

	// AccessWrite defines write access
	AccessWrite Access = "Write"

	// AccessUpdate defines update access
	AccessUpdate Access = "Update"

	// AccessNone defines void access
	AccessNone Access = "None"
)

// Validate returns true if access if valid type else returns false
func (a Access) Validate() bool {
	switch a {
	case AccessRead, AccessWrite, AccessUpdate, AccessNone:
		return true
	}

	return false
}

// CanRead returns true if access is read type, otherwise returns false
func (a Access) CanRead() bool {
	if a == AccessRead {
		return true
	}
	return false
}

// CanWrite return true if "a" has write type, otherwise returns false
func (a Access) CanWrite() bool {
	if a == AccessWrite {
		return true
	}

	return false
}

// CanUpdate returns true if "a" has update access, otherwise return false
func (a Access) CanUpdate() bool {
	if a == AccessUpdate {
		return true
	}

	return false
}

// AccessList representes a list of access
type AccessList []Access

// Validate return true/false
// Validate checks validation of all element in AccessList
func (a AccessList) Validate() bool {
	for _, v := range a {
		if !v.Validate() {
			return false
		}
	}

	return true
}

// CanRead returns true if "a" has read access
func (a AccessList) CanRead() bool {
	for _, v := range a {
		if v == AccessRead {
			return true
		}
	}

	return false
}

// CanWrite returns true if "a" has write access
func (a AccessList) CanWrite() bool {
	for _, v := range a {
		if v == AccessWrite {
			return true
		}
	}

	return false
}

// CanUpdate returns true if "a" has update access
func (a AccessList) CanUpdate() bool {
	for _, v := range a {
		if v == AccessUpdate {
			return true
		}
	}

	return false
}

// Can takes Access type and returns bool
// it checks permission according to Access
func (a AccessList) Can(acc Access) bool {
	switch acc {
	case AccessRead:
		return a.CanRead()
	case AccessWrite:
		return a.CanWrite()
	case AccessUpdate:
		return a.CanUpdate()
	}

	return false
}

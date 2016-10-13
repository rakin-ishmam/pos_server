package action

// Action is implemented by object who wants to satify action properties
type Action interface {
	Do()
	ActionErr() error
	Result() interface{}
}

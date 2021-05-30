package MRUApp

type MRU struct {
	State     State
	Item      Item
	Algorithm Algorithm
}

type State struct {
	Model string
	Item  []Item
}

type Item struct {
	label   string
	Payload Payload
}

type Payload struct {
	Content interface{}
}

type Algorithm struct {
	Logic interface{}
}

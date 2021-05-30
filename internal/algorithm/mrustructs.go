package MRUApp

type MRU struct {
	State     State     `json:"state,omitempty"`
	Item      Item      `json:"item,omitempty"`
	Algorithm Algorithm `json:"algorithm,omitempty"`
}

type State struct {
	Model string `json:"model,omitempty"`
	Item  []Item `json:"items,omitempty"`
}

type Item struct {
	Label   string  `json:"label,omitempty"`
	Payload Payload `json:"payload,omitempty"`
}

type Payload struct {
	Content interface{} `json:"content,omitempty"`
}

type Algorithm struct {
	Logic interface{} `json:"logic,omitempty"`
}

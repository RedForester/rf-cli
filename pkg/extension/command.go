package extension

type Command struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Group       *string `json:"group"`

	Type  CommandType `json:"type"`
	Rules []Rules     `json:"showRules"`
}

// TODO
type CommandType struct {
	Action string `json:"action,omitempty"`
	URL    string `json:"url,omitempty"`
}

// TODO
type Rules struct {
	AllNodes         bool   `json:"allNodes,omitempty"`
	DescendantOfType string `json:"descendantOfType,omitempty"`
	SelfType         string `json:"selfType,omitempty"`
	Root             bool   `json:"root,omitempty"`
}

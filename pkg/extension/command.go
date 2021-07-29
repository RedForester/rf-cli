package extension

type Command struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Group       *string `json:"group"`

	Type      CommandType `json:"type"`
	ShowRules []Rules     `json:"showRules" yaml:"showRules"`
}

// TODO
type CommandType struct {
	Action string `json:"action,omitempty"`
	URL    string `json:"url,omitempty"`
}

// TODO
type Rules struct {
	AllNodes         bool   `json:"allNodes,omitempty" yaml:"allNodes"`
	DescendantOfType string `json:"descendantOfType,omitempty" yaml:"descendantOfType"`
	SelfType         string `json:"selfType,omitempty" yaml:"selfType"`
	Root             bool   `json:"root,omitempty"`
}

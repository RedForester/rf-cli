package extension

type Command struct {
	Name        string
	Description string
	Group       *string

	Type  CommandType
	Rules []Rules
}

// TODO
type CommandType struct {
	Action string `json:"action,omitempty"`
	URL    string `json:"url,omitempty"`
}

// TODO
type Rules struct {
	AllNodes bool
	Root     bool
}

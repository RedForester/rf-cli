package extension

type Extension struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Email       string  `json:"email"`
	BaseURL     string  `json:"baseUrl"`
	AvatarUrl   *string `json:"avatarUrl,omitempty"`

	User          User            `json:"user"`
	Commands      []Command       `json:"commands"`
	RequiredTypes []RequiredTypes `json:"requiredTypes"`
}

type User struct {
	FirstName string  `json:"firstName,omitempty"`
	LastName  string  `json:"lastName,omitempty"`
	Username  string  `json:"username"`
	AvatarUrl *string `json:"avatarUrl,omitempty"`
}

func NewExtension() *Extension {
	return &Extension{}
}

func (e *Extension) AddCommand(cmd Command) {
	e.Commands = append(e.Commands, cmd)
}

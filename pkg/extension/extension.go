package extension

type Extension struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Email       string  `json:"email"`
	BaseURL     string  `json:"baseUrl" yaml:"baseUrl"`
	AvatarUrl   *string `json:"avatarUrl,omitempty" yaml:"avatarUrl,omitempty"`

	User          User            `json:"user"`
	Commands      []Command       `json:"commands"`
	RequiredTypes []RequiredTypes `json:"requiredTypes" yaml:"requiredTypes"`
}

type User struct {
	FirstName string  `json:"firstName,omitempty" yaml:"firstName"`
	LastName  string  `json:"lastName,omitempty" yaml:"lastName"`
	Username  string  `json:"username"`
	AvatarUrl *string `json:"avatarUrl,omitempty" yaml:"avatarUrl"`
}

func (e *Extension) AddCommand(cmd Command) {
	e.Commands = append(e.Commands, cmd)
}

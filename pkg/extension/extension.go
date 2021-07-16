package extension

type Extension struct {
	ID          string
	Name        string
	Description string
	Email       string
	BaseURL     string
	AvatarUrl   *string

	User     User
	Commands []Command
}

type User struct {
	FirstName string
	LastName  string
	Username  string
	AvatarUrl *string
}

func NewExtension() *Extension {
	return &Extension{}
}

func (e *Extension) AddCommand(cmd Command) {
	e.Commands = append(e.Commands, cmd)
}

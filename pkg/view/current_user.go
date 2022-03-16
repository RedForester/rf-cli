package view

import (
	"fmt"
	"github.com/deissh/rf-cli/pkg/rf"
	"io"
)

type CurrentUser struct {
	Data   *rf.User
	Writer io.Writer
}

func (u CurrentUser) Render() error {
	fmt.Println("ID:", u.Data.UserID)
	fmt.Println("Username:", u.Data.Username)
	fmt.Println("Name:", u.Data.Name)
	fmt.Println("Surname:", u.Data.Surname)

	return nil
}

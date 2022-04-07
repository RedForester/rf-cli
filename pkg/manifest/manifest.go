package manifest

import (
	"github.com/deissh/rf-cli/pkg/rf"
	"github.com/go-playground/validator/v10"
)

type Manifest struct {
	ID               string            `yaml:"id,omitempty" validate:"omitempty,uuid4"`
	Name             string            `yaml:"name" validate:"required"`
	Description      string            `yaml:"description"`
	ShortDescription string            `yaml:"shortDescription"`
	BaseURL          *string           `yaml:"baseUrl,omitempty"`
	Email            string            `yaml:"email" validate:"email"`
	AvatarURL        *string           `yaml:"avatarUrl,omitempty"`
	RequiredTypes    []rf.RequiredType `yaml:"requiredTypes"`
	Commands         []rf.Command      `yaml:"commands"`
	ExtensionUser    ExtUser           `yaml:"extensionUser"`
}

type ExtUser struct {
	ID        string `yaml:"id,omitempty" validate:"omitempty,uuid4"`
	Username  string `yaml:"username"`
	FirstName string `yaml:"firstName,omitempty"`
	LastName  string `yaml:"lastName,omitempty"`
	AvatarUrl string `yaml:"avatarUrl,omitempty"`
}

func FromExtension(info *rf.Extension) *Manifest {
	return &Manifest{
		ID:               info.ID,
		Name:             info.Name,
		Description:      info.Description,
		ShortDescription: info.ShortDescription,
		BaseURL:          info.BaseURL,
		Email:            info.Email,
		AvatarURL:        info.AvatarURL,
		RequiredTypes:    info.RequiredTypes,
		Commands:         info.Commands,
		ExtensionUser: ExtUser{
			ID:        info.User.ID,
			Username:  info.User.Username,
			FirstName: info.User.FirstName,
			LastName:  info.User.LastName,
			AvatarUrl: info.User.AvatarUrl,
		},
	}
}

func (m Manifest) Validate() error {
	return validator.New().Struct(m)
}

func (m Manifest) ToExtension() *rf.Extension {
	return &rf.Extension{
		ID:               m.ID,
		Name:             m.Name,
		Description:      m.Description,
		ShortDescription: m.ShortDescription,
		BaseURL:          m.BaseURL,
		Email:            m.Email,
		AvatarURL:        m.AvatarURL,
		RequiredTypes:    m.RequiredTypes,
		Commands:         m.Commands,
		User: rf.ExtUser{
			ID:        m.ExtensionUser.ID,
			Username:  m.ExtensionUser.Username,
			FirstName: m.ExtensionUser.FirstName,
			LastName:  m.ExtensionUser.LastName,
			AvatarUrl: m.ExtensionUser.AvatarUrl,
		},
	}
}

package manifest

import (
	"github.com/deissh/rf-cli/pkg/rf"
	"github.com/go-playground/validator/v10"
)

type Manifest struct {
	ID               string            `yaml:"id,omitempty" validate:"uuid4"`
	Name             string            `yaml:"name" validate:"required"`
	Description      string            `yaml:"description"`
	ShortDescription string            `yaml:"short_description"`
	BaseURL          *string           `yaml:"baseUrl,omitempty"`
	Email            string            `yaml:"email" validate:"email"`
	AvatarURL        *string           `yaml:"avatar_url,omitempty"`
	RequiredTypes    []rf.RequiredType `yaml:"required_types"`
	Commands         []rf.Command      `yaml:"commands"`
	ExtensionUser    ExtUser           `yaml:"extension_user"`
}

type ExtUser struct {
	ID        string `yaml:"id,omitempty" validate:"uuid4"`
	Username  string `yaml:"username"`
	FirstName string `yaml:"first_name,omitempty"`
	LastName  string `yaml:"last_name,omitempty"`
	AvatarUrl string `yaml:"avatar_url,omitempty"`
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

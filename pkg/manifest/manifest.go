package manifest

import (
	"github.com/deissh/rf-cli/pkg/rf"
	"github.com/go-playground/validator/v10"
)

type Manifest struct {
	ID               string            `yaml:"id,omitempty"`
	Name             string            `yaml:"name" validate:"required"`
	Description      string            `yaml:"description"`
	ShortDescription string            `yaml:"shortDescription"`
	BaseURL          *string           `yaml:"baseUrl,omitempty"`
	Email            string            `yaml:"email" validate:"email"`
	AvatarURL        *string           `yaml:"avatarUrl,omitempty"`
	RequiredTypes    []rf.RequiredType `yaml:"requiredTypes"`
	Commands         []rf.Command      `yaml:"commands"`
}

func (m Manifest) Validate() error {
	return validator.New().Struct(m)
}

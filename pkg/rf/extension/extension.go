package extension

import "encoding/json"

type Extension struct {
	ID               string         `json:"id"`
	Name             string         `json:"name"`
	Description      string         `json:"description"`
	ShortDescription string         `json:"shortDescription"`
	BaseURL          *string        `json:"baseUrl"`
	Email            string         `json:"email"`
	AvatarURL        *string        `json:"avatarUrl"`
	Published        bool           `json:"published"`
	RequiredTypes    []RequiredType `json:"requiredTypes"`
	Commands         []Command      `json:"commands"`
	User             ExtUser        `json:"user"`
	Owner            Owner          `json:"owner"`
}

type Command struct {
	Name        string      `json:"name"`
	Group       interface{} `json:"group"`
	Type        Type        `json:"type"`
	Description string      `json:"description"`
	ShowRules   []ShowRule  `json:"showRules"`
}

type Type struct {
	Action *string `json:"action"`
	URL    *string `json:"url"`
}

type Owner struct {
	ID string `json:"id"`
}

type RequiredType struct {
	Name       string     `json:"name"`
	Properties []Property `json:"properties"`
}

type ExtUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func UnmarshalExtension(data []byte) (Extension, error) {
	var r Extension
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Extension) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

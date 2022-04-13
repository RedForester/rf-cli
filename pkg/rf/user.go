package rf

import "encoding/json"

type User struct {
	UserID             string              `json:"user_id"`
	Username           string              `json:"username"`
	Name               string              `json:"name"`
	Surname            string              `json:"surname"`
	Avatar             string              `json:"avatar"`
	RegistrationDate   string              `json:"registration_date"`
	Birthday           string              `json:"birthday"`
	Contacts           Contacts            `json:"contacts"`
	IsExtensionUser    bool                `json:"is_extension_user"`
	Timezone           string              `json:"timezone"`
	Language           string              `json:"language"`
	LastAccessed       string              `json:"last_accessed"`
	Activated          bool                `json:"activated"`
	KvSession          string              `json:"kv_session"`
	CmdBuffer          []CmdBuffer         `json:"cmdBuffer"`
	SubscriptionGroups []SubscriptionGroup `json:"subscription_groups"`
	Tags               []Tag               `json:"tags"`
	SavedSearchQueries []interface{}       `json:"saved_search_queries"`
}

type CmdBuffer struct {
	ID      string        `json:"id"`
	Meta    CmdBufferMeta `json:"meta"`
	Type    string        `json:"type"`
	Nodes   []string      `json:"nodes"`
	Branch  bool          `json:"branch"`
	Oneshot bool          `json:"oneshot"`
}

type CmdBufferMeta struct {
	Map struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"map"`
	Titles []string `json:"titles"`
}

type Contacts struct {
}

type SubscriptionGroup struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	UserCanChange  bool   `json:"user_can_change"`
	WillBeNotified bool   `json:"will_be_notified"`
}

type Tag struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Removable bool   `json:"removable"`
}

func UnmarshalUser(data []byte) (User, error) {
	var r User
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *User) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

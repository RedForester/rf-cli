package rf

import "encoding/json"

type Map struct {
	ID            string `json:"id"`
	RootNodeID    string `json:"root_node_id"`
	Owner         string `json:"owner"`
	OwnerName     string `json:"owner_name"`
	OwnerSurname  string `json:"owner_surname"`
	OwnerUsername string `json:"owner_username"`
	OwnerAvatar   string `json:"owner_avatar"`
	Public        bool   `json:"public"`
	NodeCount     int64  `json:"node_count"`
	UserCount     int64  `json:"user_count"`
	Name          string `json:"name"`
	IsAdmin       bool   `json:"is_admin"`
	CanExport     bool   `json:"can_export"`
	Description   string `json:"description"`
}

func UnmarshalMap(data []byte) (Map, error) {
	var r Map
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Map) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

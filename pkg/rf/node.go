package rf

import "encoding/json"

func UnmarshalNode(data []byte) (Node, error) {
	var r Node
	err := json.Unmarshal(data, &r)
	return r, err
}

func (n *Node) Marshal() ([]byte, error) {
	return json.Marshal(n)
}

type Node struct {
	ID                  string         `json:"id"`
	MapID               string         `json:"map_id"`
	Parent              string         `json:"parent"`
	Position            []Position     `json:"position"`
	Properties          NodeProperties `json:"properties"`
	UnreadCommentsCount int64          `json:"unread_comments_count"`
	CommentsCount       int64          `json:"comments_count"`
	Access              string         `json:"access"`
	OriginalParent      string         `json:"originalParent"`
	Body                Body           `json:"body"`
	Hidden              bool           `json:"hidden"`
	Readers             []string       `json:"readers"`
	Nodelevel           int64          `json:"nodelevel"`
	Meta                Meta           `json:"meta"`
}

func (n Node) IsLink() bool {
	return n.ID != n.Body.ID
}

type Body struct {
	ID                  string         `json:"id"`
	MapID               string         `json:"map_id"`
	TypeID              string         `json:"type_id"`
	Properties          NodeProperties `json:"properties"`
	Parent              string         `json:"parent"`
	UnreadCommentsCount int64          `json:"unread_comments_count"`
	CommentsCount       int64          `json:"comments_count"`
	Children            []Node         `json:"children"`
	Access              string         `json:"access"`
	Readers             []string       `json:"readers"`
	Meta                Meta           `json:"meta"`
}

type Meta struct {
	CreationTimestamp     string `json:"creation_timestamp"`
	LastModifiedTimestamp string `json:"last_modified_timestamp"`
	LastModifiedUser      string `json:"last_modified_user"`
	Author                string `json:"author"`
	Subscribed            *bool  `json:"subscribed,omitempty"`
	CanMove               bool   `json:"can_move"`
	Editable              bool   `json:"editable"`
	Commentable           bool   `json:"commentable"`
	CanSetAccess          bool   `json:"can_set_access"`
	Leaf                  *bool  `json:"leaf,omitempty"`
}

type NodeProperties struct {
	Style       map[string]string      `json:"style"`
	ByType      map[string]interface{} `json:"byType"`
	ByUser      []PropByUser           `json:"byUser"`
	Global      GlobalProperties       `json:"global"`
	ByExtension map[string]interface{} `json:"byExtension"`
}

type PropByUser struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	TypeID  int64  `json:"type_id"`
	Visible bool   `json:"visible"`
}

type GlobalProperties struct {
	Title string `json:"title"`
}

type Position struct {
	Integer *int64
	String  *string
}

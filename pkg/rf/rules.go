package rf

type ShowRule struct {
	SelfType         *string `json:"selfType,omitempty"`
	DescendantOfType *string `json:"descendantOfType,omitempty"`
	AllNodes         bool    `json:"allNodes,omitempty"`
	Root             bool    `json:"root,omitempty"`
}

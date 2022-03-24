package rf

type ShowRule struct {
	SelfType         *string `json:"selfType,omitempty" yaml:"selfType,omitempty"`
	DescendantOfType *string `json:"descendantOfType,omitempty" yaml:"descendantOfType,omitempty"`
	AllNodes         bool    `json:"allNodes,omitempty" yaml:"allNodes,omitempty"`
	Root             bool    `json:"root,omitempty" yaml:"root,omitempty"`
}

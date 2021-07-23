package extension

type RequiredTypes struct {
	Name       string     `json:"name"`
	Properties []Property `json:"properties"`
}

type Property struct {
	Argument string `json:"argument"`
	Category string `json:"category"`
	Name     string `json:"name"`
}

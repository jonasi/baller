package spec

type Endpoint struct {
	Name       string      `json:"name"`
	Path       string      `json:"path"`
	Parameters []Value     `json:"parameters"`
	ResultSets []ResultSet `json:"result_sets"`
}

type Value struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ResultSet struct {
	Name   string  `json:"name"`
	Values []Value `json:"values"`
}

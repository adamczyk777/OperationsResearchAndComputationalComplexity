package main

type SubAhp struct {
	Name        string `json:"name"`
	Preferences [][]float64 `json:"preferences"`
	Children    []SubAhp `json:"children"`
}

type Ahp struct {
	Alternatives []string `json:"alternatives"`
	Goal         SubAhp `json:"goal"`
}

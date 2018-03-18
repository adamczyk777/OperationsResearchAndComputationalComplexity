package main

type SubAhp struct {
	Name        string
	Preferences [][]float64
	Children    []SubAhp
}

type Ahp struct {
	Alternatives []string
	Goal         SubAhp
}

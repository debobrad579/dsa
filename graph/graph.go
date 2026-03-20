package graph

type AdjacencyMatrix [][]int

type AdjacencyList [][]Edge

type Edge struct {
	To     int
	Weight int
}

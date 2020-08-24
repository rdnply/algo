package main

import "fmt"

const (
	White = iota
	Gray
	Black
)

type graph struct {
	v   int
	adj [][]int
}

func NewGraph(v int) *graph {
	adj := make([][]int, v)
	for i := 0; i < v; i++ {
		adj[i] = make([]int, 0)
	}

	return &graph{v, adj}
}

func (g *graph) addEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
}

func (g graph) isCyclicUtil(v int, colors map[int]int) bool {
	colors[v] = Gray

	for _, to := range g.adj[v] {
		if colors[to] == Gray {
			return true
		}
		if colors[to] == White && g.isCyclicUtil(to, colors) {
			return true
		}
	}

	colors[v] = Black
	return false
}

func (g graph) isCyclic() bool {
	colors := make(map[int]int)

	for i := 0; i < g.v; i++ {
		colors[i] = White
	}

	for i := 0; i < g.v; i++ {
		if colors[i] == White {
			if g.isCyclicUtil(i, colors) {
				return true
			}
		}
	}

	return false
}

func main() {
	g := NewGraph(5)
	g.addEdge(1, 0)
	g.addEdge(0, 2)
	g.addEdge(2, 1)
	g.addEdge(0, 3)
	g.addEdge(3, 4)
	fmt.Println(g.isCyclic())

	g2 := NewGraph(3)
	g2.addEdge(0, 1)
	g2.addEdge(1, 2)
	fmt.Println(g2.isCyclic())
}

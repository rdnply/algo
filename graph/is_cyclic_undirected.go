package main

import "fmt"

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

func (g graph) isCyclic() bool {
	visited := make(map[int]bool)

	for i := 0; i < g.v; i++ {
		if !visited[i] {
			if g.isCyclicUtil(i, visited, -1) {
				return true
			}
		}
	}

	return false
}

func (g graph) isCyclicUtil(v int, visited map[int]bool, parent int) bool {
	visited[v] = true

	for _, to := range g.adj[v] {
		if !visited[to] {
			if g.isCyclicUtil(to, visited, v) {
				return true
			}
		} else if to != parent {
			return true
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

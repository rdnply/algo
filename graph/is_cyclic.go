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

func (g graph) isCyclicUtil(v int, visited, recStack map[int]bool) bool {
	if !visited[v] {
		visited[v], recStack[v] = true, true

		for _, to := range g.adj[v] {
			if !visited[to] && g.isCyclicUtil(to, visited, recStack) {
				return true
			} else if recStack[to] {
				return true
			}
		}
	}

	recStack[v] = false
	return false
}

func (g graph) isCyclic() bool {
	visited := make(map[int]bool)
	recStack := make(map[int]bool)

	for i := 0; i < g.v; i++ {
		if g.isCyclicUtil(i, visited, recStack) {
			return true
		}
	}

	return false
}

func main() {
	g := NewGraph(4)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 0)
	g.addEdge(2, 3)
	g.addEdge(3, 3)

	fmt.Println(g.isCyclic())
}

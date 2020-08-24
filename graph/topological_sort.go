package main

import (
	"fmt"

	"../stack/stack"
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

func (g graph) topologicalSort() []int {
	visited := make(map[int]bool)
	st := stack.New()

	for i := 0; i < g.v; i++ {
		if !visited[i] {
			g.topSortUtil(i, visited, &st)
		}
	}

	res := make([]int, 0)
	for !st.Empty() {
		res = append(res, st.Pop())
	}

	return res
}

func (g graph) topSortUtil(v int, visited map[int]bool, st *stack.Stack) {
	visited[v] = true

	for _, to := range g.adj[v] {
		if !visited[to] {
			g.topSortUtil(to, visited, st)
		}
	}

	st.Push(v)
}

func main() {
	g := NewGraph(6)
	g.addEdge(5, 2)
	g.addEdge(5, 0)
	g.addEdge(4, 0)
	g.addEdge(4, 1)
	g.addEdge(2, 3)
	g.addEdge(3, 1)

	res := g.topologicalSort()

	fmt.Println(res)
}

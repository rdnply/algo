package main

import (
	"fmt"

	"../queue/queue"
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

func (g graph) BFS(s int) []int {
	res := make([]int, 0)
	visited := make(map[int]struct{})

	q := queue.New()
	q.Enqueue(s)
	visited[s] = struct{}{}

	for !q.Empty() {
		t := q.Dequeue()
		res = append(res, t)

		for _, v := range g.adj[t] {
			if _, ok := visited[v]; !ok {
				visited[v] = struct{}{}
				q.Enqueue(v)
			}
		}
	}

	return res
}

func (g graph) DFS(s int) {
	visited := make(map[int]struct{})

	g.DFSUtil(s, visited)
	fmt.Println()
}

func (g graph) DFSUtil(v int, visited map[int]struct{}) {
	fmt.Print(v, " ")
	visited[v] = struct{}{}
	for _, to := range g.adj[v] {
		if _, ok := visited[to]; !ok {
			g.DFSUtil(to, visited)
		}
	}
}

// func (g graph) DFSIter(s int) []int {
//     res := make([]int, 0)
//     visited := make(map[int]struct{})
//
//     st := stack.New()
//     st.Push(s)
//
//     for !st.Empty() {
//         v := st.Pop()
//
//         if _, ok := visited[v]; !ok {
//             visited[v] = struct{}{}
//             res = append(res, v)
//         }
//
//         for _, to := range g.adj[v] {
//             if _, ok := visited[to]; !ok {
//                 st.Push(to)
//             }
//         }
//     }
//
//     return res
// }

func main() {
	g := NewGraph(5)
	g.addEdge(1, 0)
	g.addEdge(0, 2)
	g.addEdge(2, 1)
	g.addEdge(0, 3)
	g.addEdge(1, 4)

	res := g.BFS(0)
	fmt.Println(res)

	g.DFS(2)

	// res = g.DFSIter(2)
	// fmt.Println(res)
}

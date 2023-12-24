package graph

import (
	"fmt"
)

type Edge struct {
	Source int `json:"source"`
	Drain  int `json:"drain"`
}

type Graph struct {
	VertexCount int    `json:"vertexesCount"`
	Edges       []Edge `json:"edges"`
	vertexes    [][]int
}

func (g *Graph) Init() {
	g.vertexes = make([][]int, g.VertexCount)

	for _, edge := range g.Edges {
		g.vertexes[edge.Source] = append(g.vertexes[edge.Source], edge.Drain)
	}
}

func (g *Graph) FindPath(startId int, finishId int) (ans []int, err error) {
	visited := make([]bool, g.VertexCount)
	found := false

	var rec func(v int)
	rec = func(v int) {
		visited[v] = true
		ans = append(ans, v)

		if v == finishId {
			found = true
			return
		}

		for _, u := range g.vertexes[v] {
			if !visited[u] && !found {
				rec(u)
			}
		}
		if !found {
			ans = ans[:len(ans)-1]
		}
	}

	rec(startId)

	if !found {
		err = fmt.Errorf("path not found")
	}

	return
}

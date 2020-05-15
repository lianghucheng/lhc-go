package main

import "fmt"

func main() {
	adjacencyMatrix := [][]int{
		{0,1,0,0,0,0,0,0,0,0},
		{1,0,0,1,0,1,0,1,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,1,0,1,0,1,0,1,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,1,0,1,0,1,0,1,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,1,0,1,0,1,0,1,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
	}

	visit := make(map[int]bool)
	ret := recuiceAdjacencyMatrixDfs(adjacencyMatrix, visit, []int{0})
	fmt.Println(ret)
}

func adjacencyMatrixDfs1 (graph [][]int, startNode int) []int {
	result := []int{}
	visit := make(map[int]bool)
	stack := []int{startNode}

	for ;len(stack) > 0; {
		s := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if visit[s] {
			continue
		}
		visit[s] = true
		result = append(result, s)
		for k,v := range graph[s] {
			if v == 1 && !visit[k] {
				stack = append(stack, k)
			}
		}
	}

	return result
}

func recuiceAdjacencyMatrixDfs(graph [][]int, visit map[int]bool, stack []int) []int {
	if len(stack) == 0 {
		return []int{}
	}
	node := stack[len(stack) - 1]
	stack = stack[:len(stack) - 1]
	result := []int{node}
	if visit[node] == true {
		return []int{}
	}
	visit[node] = true
	for k, v := range graph[node] {
		if !visit[k] && v == 1 {
			stack = append(stack, k)
		}
	}

	result = append(result,recuiceAdjacencyMatrixDfs(graph, visit, stack)...)

	return result
}

func recursionAdjacencyMatrixDfsByDepth(graph [][]int, visit map[int]bool, stack []int) (result []int) {
	return
}
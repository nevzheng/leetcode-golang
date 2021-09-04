package _834_sum_of_distances_in_tree

func makeAdjList(n int, edges [][]int) [][]int {
	adjList := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}
	return adjList
}

func dfs(adjList [][]int, i, p int, count, dist []int) {
	for _, child := range adjList[i] {
		if child != p {
			dfs(adjList, child, i, count, dist)
			count[i] += count[child] // Sum up all child counts
			dist[i] += dist[child] + count[child]
		}
	}
	count[i] += 1 // count node i, the root of the dfs tree
}

func dfs2(adjList [][]int, i, p, n int, count, dist []int) {
	for _, child := range adjList[i] {
		if child != p {
			dist[child] = dist[i] - count[child] + n - count[child]
			dfs2(adjList, child, i, n, count, dist)
		}
	}
}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	adjList := makeAdjList(n, edges)
	count := make([]int, n) // Number of nodes in subtree rooted at node n
	dist := make([]int, n)
	dfs(adjList, 0, -1, count, dist)     // populate the count and dist arrays
	dfs2(adjList, 0, -1, n, count, dist) // recursively pivot, re-root the answer
	return dist
}

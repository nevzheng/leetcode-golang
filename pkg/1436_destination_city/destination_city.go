package _436_destination_city

func destCity(paths [][]string) string {
	outDegree := make(map[string]int) // Count Outgoing edges for each city
	for _, path := range paths {
		// Ensure the both cities are in the map
		u, v := path[0], path[1]
		if _, ok := outDegree[u]; !ok {
			outDegree[u] = 0
		}
		if _, ok := outDegree[v]; !ok {
			outDegree[v] = 0
		}
		outDegree[u]++ // Increment outDegree
	}
	// Return a City with no outgoing edges
	for s, i := range outDegree {
		if i == 0 {
			return s
		}
	}
	return "" // No Destination Cities Exist
}

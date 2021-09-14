package main

var judgePairs = [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}

func judge() string {
	for _, v := range judgePairs {
		if state[v[0]] == state[v[1]] && state[v[1]] == state[v[2]] {
			return state[v[0]]
		}
	}
	return ""
}

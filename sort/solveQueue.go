package main

// TODO
func solveQueue(n int) [][]string {
	s := make([][]string, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

		}
	}

}

func getFreeQueue(row, col, n int) map[int]int {
	mapper := make(map[int]int)
	v := row - col
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == row || j == col {
				continue
			}
			if i-j == v || j-i == v {
				continue
			}
			mapper[i] = j
		}
	}
	return mapper
}

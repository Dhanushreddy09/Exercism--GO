package chessboard

type Rank []bool
type Chessboard map[string]Rank

func CountInRank(cb Chessboard, rank string) int {
	rankArray, rankExists := cb[rank]
	count := 0

	if !rankExists {
		return 0
	}

	for _, val := range rankArray {
		if val {
			count++
		}
	}
	return count
}

func CountInFile(cb Chessboard, file int) int {
	count := 0
	if file < 1 || file > 8 {
		return 0
	}
	for _, val := range cb {
		if val[file-1] {
			count++
		}
	}
	return count
}

func CountAll(cb Chessboard) int {
	length := len(cb)
	return length * length
}

func CountOccupied(cb Chessboard) int {
	count := 0
	for key := range cb {
		count += CountInRank(cb, key)
	}
	return count
}

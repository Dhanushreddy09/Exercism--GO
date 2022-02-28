package chessboard

import "fmt"

func Chess() {
	cb := make(map[string]Rank)
	cb["a"] = Rank{true, false, true, true, false, true, true, true}
	cb["b"] = Rank{true, false, false, true, true, false, true, false}
	cb["c"] = Rank{false, true, true, false, true, true, true, true}
	cb["d"] = Rank{true, false, true, true, true, true, true, true}
	cb["e"] = Rank{true, true, true, true, true, true, true, true}
	cb["f"] = Rank{true, true, true, true, true, true, true, true}
	cb["g"] = Rank{true, true, true, true, true, true, true, true}
	cb["h"] = Rank{true, true, true, true, true, true, true, true}

	fmt.Println(CountInRank(cb, "a"))
	fmt.Println(CountInFile(cb, 1))
	fmt.Println(CountAll(cb))
	fmt.Println(CountOccupied(cb))
}

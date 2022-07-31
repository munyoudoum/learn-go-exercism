package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
// type Rank struct {
// 	make([]bool)
// }
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank string) int {
	res := 0
	for _, r := range cb[rank] {
		if r {
			res += 1
		}
	}
	return res
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file int) int {
	// fmt.Println(cb, file)
	if file < 1 || file > 8 {
		return 0
	}
	res := 0
	for _, r := range cb {
		if r[file-1] {
			res += 1
		}
	}
	return res
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	res := 0
	for _, r := range cb {
		res += len(r)

	}
	return res
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	res := 0
	for _, r := range cb {
		for _, file := range r {
			if file {
				res += 1
			}
		}
	}
	return res
}

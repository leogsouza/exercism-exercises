package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	c := 0
	squares, ok := cb[file]
	if !ok {
		return c
	}

	for _, v := range squares {
		if v {
			c++
		}
	}

	return c
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	c := 0
	if rank > 8 || rank < 1 {
		return c
	}
	idxRank := rank - 1

	for _, v := range cb {
		if v[idxRank] {
			c++
		}
	}
	return c
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	sum := 0
	for _, v := range cb {
		sum += len(v)
	}
	return sum
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	c := 0
	for _, f := range cb {
		for _, s := range f {
			if s {
				c++
			}
		}
	}
	return c
}

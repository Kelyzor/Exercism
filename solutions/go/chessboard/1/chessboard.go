package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    k := 0
	for _, x := range cb[file] {
        if x == true {
            k++
        }
    }
    return k
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    if rank > 8 || rank < 1 {
        return 0
    }
    k := 0
    for _, x := range cb {
        if x[rank - 1] == true {
            k++
        }
    }
    return k
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    k := 0
    for _, x := range cb {
        for _, y := range x {
            if y || !y {
                k++
            }
        }
    }
    return k
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	k := 0
    for _, x := range cb {
        for _, y := range x {
            if y {
                k++
            }
        }
    }
    return k
}

package main

import "fmt"

type directions struct {
	index int
	base  string
}

func (d *directions) next() string {

	if d.index == len(d.base) {
		d.index = 0
	}
	s := string(d.base[d.index])
	d.index++
	return s
}

type piece struct {
	index  int
	pieces [][][]string
}

func (p *piece) next() [][]string {
	if p.index == len(p.pieces) {
		p.index = 0
	}
	mypiece := p.pieces[p.index]
	p.index++
	return mypiece
}

type board struct {
	columns       [][]string
	_cache        map[MyKey]MyValue
	repeatsHeight int
	foundLoop     bool
}

func (b *board) print() {
	fmt.Printf("========================\n")
	for i := len(b.columns) - 1; i >= 0; i-- {
		fmt.Printf("| ")
		for j := 0; j < len((b.columns)[i]); j++ {
			fmt.Printf("%s ", (b.columns)[i][j])
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("========================\n")
}

func (b *board) highestLevel() int {
	for i := len(b.columns) - 1; i >= 0; i-- {
		if b.lineIsEmpty(i) == false {
			return i
		}
	}
	return 0
}

func (b *board) canMovePieceRight() bool {
	for i := 0; i < len(b.columns); i++ {
		for j := 0; j < len(b.columns[i]); j++ {
			if b.columns[i][j] == "@" {
				if j == len(b.columns[i])-1 {
					return false
				}
				if b.columns[i][j+1] == "#" {
					return false
				}
			}
		}
	}
	return true
}

func (b *board) getStringRepresentationTopXRows(n int) string {
	var s string = ""
	for r := len(b.columns) - 1; r >= len(b.columns)-n; r-- {
		for j := 0; j < len(b.columns[r]); j++ {
			s += b.columns[r][j]
		}
	}
	// fmt.Printf("String repre for top %d rows: %s\n", n, s)
	return s
}

type MyValue struct {
	height     int
	pieceCount int
}

type MyKey struct {
	pieceIndex   int
	lateralIndex int
	topRows      string
}

func (b *board) _check_cache(pieceIndex int, lateralIndex int, topRows string, pieceCount int) (bool, int, int, int, int) {
	key := MyKey{
		pieceIndex:   pieceIndex,
		lateralIndex: lateralIndex,
		topRows:      topRows,
	}
	currentHeight := b.highestLevel()
	if val, ok := b._cache[key]; ok {
		return true, currentHeight, val.height, pieceCount, val.pieceCount

	} else {
		val = MyValue{
			height:     currentHeight,
			pieceCount: pieceCount,
		}
		b._cache[key] = val
	}
	return false, currentHeight, 0, pieceCount, 0

}

func (b *board) calculateHeight(totalPieces int, currentPieces int, piecesDelta int, heightDelta int) (int, int) {
	remainingPieces := totalPieces - currentPieces
	requiredRepeats := remainingPieces / piecesDelta
	remainingPiecesAfterRepetition := remainingPieces % piecesDelta
	repeatAddedHeight := heightDelta * requiredRepeats
	return repeatAddedHeight, remainingPiecesAfterRepetition
}

func (b *board) movePieceRight() {
	for i := 0; i < len(b.columns); i++ {
		for j := len(b.columns[i]) - 2; j >= 0; j-- {
			if b.columns[i][j] == "@" {
				b.columns[i][j+1] = "@"
				b.columns[i][j] = "."
			}
		}
	}
}

func (b *board) canMovePieceLeft() bool {
	for i := 0; i < len(b.columns); i++ {
		for j := 0; j < len(b.columns[i]); j++ {
			if b.columns[i][j] == "@" {
				if j == 0 {
					return false
				}
				if b.columns[i][j-1] == "#" {
					return false
				}
			}
		}
	}
	return true
}

func (b *board) movePieceLeft() {
	for i := 0; i < len(b.columns); i++ {
		for j := 1; j < len(b.columns[i]); j++ {
			if b.columns[i][j] == "@" {
				b.columns[i][j-1] = "@"
				b.columns[i][j] = "."
			}
		}
	}
}

func (b *board) movePieceLaterally(move string) {
	if move == ">" {
		// fmt.Println("Moving piece to the right")
		if b.canMovePieceRight() {
			// fmt.Println("Piece moved to the right")
			b.movePieceRight()
		} else {
			// fmt.Println("Could not move piece to the right")
		}
	} else if move == "<" {
		// fmt.Println("Moving piece to the left")
		if b.canMovePieceLeft() {
			// fmt.Println("Piece moved to the left")
			b.movePieceLeft()
		} else {
			// fmt.Println("Could not move piece to the left")
		}
	}
}

func (b *board) appendEmptyLine() {
	line := make([]string, 0)
	for i := 0; i < 7; i++ {
		line = append(line, ".")
	}
	b.columns = append(b.columns, line)
}

func (b *board) putPieceInBoard(p [][]string) {
	newPiece := make([][]string, 0)
	for _, line := range p {
		newLine := make([]string, len(line))
		copy(newLine, line)
		newPiece = append(newPiece, newLine)
	}
	// fmt.Printf("Putting this piece %+v\n", p)
	for i := 0; i < len(newPiece); i++ {
		b.columns = append(b.columns, newPiece[i])
	}
}

func (b *board) canPieceMoveDown() bool {
	for i := 0; i < len(b.columns); i++ {
		for j := 0; j < len(b.columns[i]); j++ {
			if b.columns[i][j] == "@" {
				if i == 0 {
					return false
				}
				if b.columns[i-1][j] == "#" {
					return false
				}

			}
		}
	}
	return true
}

func (b *board) movePieceDown() {
	for i := 1; i < len(b.columns); i++ {
		for j := 0; j < len(b.columns[i]); j++ {
			if b.columns[i][j] == "@" {
				if b.columns[i-1][j] != "#" {
					b.columns[i-1][j] = "@"
					b.columns[i][j] = "."
				} else {
					fmt.Printf("SHOULD NOT HAPPEN \n")
					return
				}
			}
		}
	}
}

func (b *board) settlePiece() {
	for i := 0; i < len(b.columns); i++ {
		for j := 0; j < len(b.columns[i]); j++ {
			if b.columns[i][j] == "@" {
				b.columns[i][j] = "#"
			}
		}
	}
}

func (b *board) prepareBoardForPieceStart() {
	i := 0
	fmt.Printf("Highest level %d\n", b.highestLevel())
	for i = b.highestLevel(); i < len(b.columns); i++ {
		if b.lineIsEmpty(i) != false {
			break
		}
	}
	if len(b.columns) > i+3 {
		for len(b.columns) > i+3 {
			b.columns = b.columns[:len(b.columns)-1]
		}
	}
	if len(b.columns) < i+3 {
		for len(b.columns) < i+3 {
			b.appendEmptyLine()
		}
	}
}

func (b *board) lineIsEmpty(i int) bool {
	empty := true
	for j := 0; j < len((b.columns)[i]); j++ {
		if (b.columns)[i][j] != "." {
			empty = false
			break
		}
	}
	return empty
}

func printPiece(p [][]string) {
	for i := len(p) - 1; i >= 0; i-- {
		for j := 0; j < len(p[i]); j++ {
			fmt.Printf(" %s", p[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {
	TOTAL_PIECES := 1000000000000
	d := &directions{
		index: 0,
		// base:  ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>",
		base: ">><>>><<<>>><<>>>><<<>>>><<><>>><<<<>><<<><<<>><<<<>><<>>>><<>><<<><>><><<<>><<<<>>>><<>>><><<>>><<><<<<>>><<>>><>><<<<>><<<<>><<><<>>>><<<>>><<>>>><<<><><><<>>>><<>><<>><>>><<<>><<>>><<<<>>><><<<><<<<><<<<><<<<>>>><>>>><<<>>>><>>><>>><<<<>>><<<<><<<<>>><>>>><<>>>><<<<>>>><<<><<><<<>>>><>>><<<<>>><<<><<>>><<<<>>>><<<>><>><>><<<<>>>><><>>>><<<<>>>><<><<<<><<<<>>>><<>>><<>><<<><<<><<<<>>><<<>>><<>>>><<<>>>><>>>><<<>><<<>>><<<<>>><<>>><<<<>><<>>>><><<<<>>><><>><>><<<>>><<<>><>>>><<<<>>><<<<>>><<<<>><<<<>>><<>>>><>>>><<<>>><<<<><<<>><><<<<>>>><<<>><<<<>>>><<>><<<><<<<><<<<>>>><<<>>>><<>><><<<>><<>>>><<<>><<<>>>><>>>><><<<<><><<>>><<<>><<>>>><<<>><<>>>><<>>><<<<><<<><<>><<<>><>><<<>>>><><>>>><<<><>>>><<<>>>><<>><<>><<<<><<<>>><<<>><<<<>>><>><<>>><<<<>>>><<<><<><<>><<<<>>><<><<<<>>><>>><<>>><<>><<><><<>>>><<><<<<>>><<<<>>><<<<>>>><<>>>><>><<<<>><<<<>>><>><<<>>>><<><<<>>>><<>>><<>><<<<>><<<><<<<>>><<><<<<>>>><<>>><<<>>>><<<>><<<>>>><>>>><>>><>>><>>><>>><<>>><<><<<<>><<>><<>><>>>><<<<>>><<><<>>>><<<><<>>><<><<<>>><<>><>><<><<<>>><<>>>><<<<>><<><<>><>>>><><<<<>>>><<><<><<>><<<<>><<<><<>>>><<<<>>><>><<>>><>><<<<>><<>>><><<<>>><<<>>>><>>><<<>>><><<<><<<><<<>><>>>><<><<>>><<<<>>>><<>>><<>>>><<<<>>><<>><<<<>><<<<>>><><<><<<>><<>>><><>>>><<<><>>>><<>>>><<<>>><<>><>>>><>><<>>>><<<<>><<<>><<>><>>><<<<>>>><><<<<>>><<<<>>><>>><<<<>>>><>>><<<<>>>><><<>>><<><<<>>><<<>>><>>><<>>>><><<<<>><<>>>><<>>>><<>>><<>><<><>>>><<>>>><<<<>>><<<<><<<<><<>>><<<>><<<>>>><<><<<<>>>><<<<>>>><<<>>><><<<>>><<<>><<<>><<<<>>><<>>><<>><<<>>>><<<<>><><><<<>>>><<<>>><<>>>><<>>>><<<><><<<>>>><<<>>><>><<<>>>><<<<><<<<>>>><<<>>><<>><>>><<>><>>>><<>>><<>><<<>>><><<<>><<<>><><><<<<><<<<><<<<>><>>>><<>>><>><<<>>>><<>><<<>>><<>><<<>>><>>><<<>>>><>><<>><<<>><>><<>>>><>>>><<<><<<<>><>>>><>>>><>>><<<><<>>>><<>>>><<<<><<>>>><<<<><<>>>><<>>><<>><<<>><<>>><><<<<><>>>><>>><<<>><<<<><>>><<<<>>><>>>><><<>>>><<<<>>>><>><>>>><<<<><<<<>>>><>>>><<<>>><><<><<>>>><<<>><<<<><<<<>>><<>>><<<<>>>><<<<>><><<>>><>><<<>>><<>><<>>><><<>><<<<>>>><<<><<<<>>>><<><<<>>>><<<<><<>>>><>>>><<<<>>>><>>>><<<>><><<<>><<<<><<<>><>>><<<<>><<<<>>><<>><>>><><<<<>>><<><<<<>>>><>>><>><<<<><<<>>><>><<<>><<<><<>><>><<<<>>><<<>>><><>>>><><>><<<>>><<><<>>>><<<>>>><><>>><<<><<>><><<<><<>>>><<<>><<><>><<<>><<<<>>><<>>><<>><<<>><><<>>><<>>>><<<>><>><<>>>><<<>>>><<>>><<>>>><<>>><>><>><<<>>><<>>>><>><<>>>><<<<>><<<><<<<>><>>><<>>><<<<><<<>><><<<<>><<><<<>><<<>><>>>><<><>><<<<>>>><>>>><<><><<<<>>>><<<>>>><<<>><<><<<<>>>><<<<>><<<>>>><<<<>>>><<>><<<<>>>><>>><<><<<>>>><<>>><>>>><>><>><>><<>>>><<<><>><<<>>>><<<<>>>><<<<><<<>><><<><<>>>><<<<>><<<<>><><<<<>>>><<<<>><><<>>>><<>><<>>>><<><<><<<><<>>><<><<<><<><<<>>>><<<<>>><><<<<><<>>>><>><<<>>>><<>>><<>>><>>><>>>><<<>>><<>><<>>>><<<>>>><<<<>><<<>><<<><<<><><<<<>>><<>>><<<>><<>>>><><<<<>><<<<>>>><<<<>>><<<<><<>><<>>><<<<><<<><<<<>><<<>><<<><<<>><<<>><<<<>>>><<<>>><>>><<<<>><<<>><<<>>><>>>><<<<>><<<<><<><>><>><<><<<>>>><<<<><<>>><><<<<>><<<>>>><<<<>>><<>>>><<>>>><<<>>><<>>>><<>><>>><<><<<<>>><<<<>><<<<>><<<<>>><<<<><<><><<<><>><<<<>><<<>>>><<<<>>>><<<<>>><<><<><>><<>><<><<<<><<><>><>><<<>><<>>><>>>><<<><<<>><<<>>>><<>>><<>>>><<<<>>><<<<>><<<<>>>><<>><>>><><>><<<<>>>><<<>><<>>><<<>>>><>>>><<<<>>><<>>>><<<><<<<>>>><<><>>>><<<>>>><<><<<<>><<<<>><<><<<<>>>><<<>>>><<<<>><<>><>><<><<>>>><<>>>><<><<>><<<<>>><<<<>>>><>>><<<><>>>><<<>><<<>>><<>>><<<<>>>><><<<<>><>><<>>><<<>>>><<<<>><><>><<<<><<<<>>><<<<>>>><<<>>><<<<>>>><<>>>><<<><<>>><<>>><<<><>><>>>><<<<>>><<<<>>>><<<>><<>>><>>>><<<<>>><<>>>><<<>>>><<<><<<<>><>>><<><><<<>><<<>>><<><<<<>>>><<<<><<<<><<<>>>><<<<><<>><<<<>><<>>>><>><<<<><>>><<<><<<><<>>>><>><<<><<<<><<>>>><<>><<<<>>>><<>><<<<>><<<>><<><>>><>><<><<>>><<><<<<>><><<>>><<<>>><>>><<>>><<<><<<>>>><<<<>>><<<<>>>><<<<>>><<<<>>><<<>>>><<<>><<>><<<<>>>><<<<><<>>>><<>>><<<><<><<>>><<<<>>><<<<><<<>>><>>><<<>>>><>>><<>>><<>>>><>>><<<>>><<<><<<<>><<><<<>>><<<>>><<>>>><<><<>>><<><<<<>>>><<<>><<<<>>><<<<>>><>>>><><<<>><<<<>>>><>><<<>>><<>>>><<<<><<<>>><<<>>><><>>>><>><<<>>><<<<><<<>><<<<>>>><<<<><><<>><><>><<<><<<<><<<>>><<>>>><<<>>><<>>><<>>><<>>><<<>>>><<><<>>>><<<<>>>><<<>><<<><<<<>>><<>>>><<<<>>>><<><<>>><<>><<<>>>><<<>>>><>><<<<>>>><>>>><>>><<<<>><<<<>>><<<><<<>>><<<>><<<>><<<><>>><>><<<<>><<<<>>>><<>><<>><<<>>>><<<>><<<<>>><<>>>><<<<><><<<>>><>>><<<<>>><<>>><<<<>><<<>>><<<<>><>>>><<>>><><<<>>>><<<<>>>><>><<<>>>><<<>>>><>><>>><<<>>>><><>>>><<<<><>>><<<>><<<<>>><<<>>><<><<>>>><<<><<<><<<>>><<<><>><>>><><<<<>>><<<<>>><>>>><<<<>>>><>>><<<>>><>>><>>><<<><<><<><<<<><<<<>>>><<<>>>><<<>>>><<<<><<<><<>>>><>>>><<<>><<<>>><<>><<<<><><>><<><<<>>><<<<>>>><><<><<<>><>>>><<<<>>>><<>>><<<<><<<>>>><<>>><<<<>>>><<<<>><<>>><<<>>>><<<<>><<<><<>><>><<>><>><<>>><<>>>><<<<>><<><<<><<<<>><<<<>><>>><><<<<>><<>><<<<>><<<>>><><<<><<<>>><<><<<<>><<<<>>><<><<>><><<<<><<<>><<<><<<<>>><<<<><<>>><<<>>><<<>><<<<><<>>>><<<<>>><<<<>>><<<<><>>>><<><>><>>><<>>>><><<<><><>>><<<<><<><<<>>><>>><><<<><<>><>>><<<<>><<<>>><<>>>><<>>>><>>>><<<>>><>><<<<>><<<>>><<<>>>><>>><<><<><<<<>>>><<<><>>>><<>>>><<<>>>><<>>>><>><<<<>>>><<><<<<>>>><<>><<<>>>><>><<><<<<>><>>>><<<<>><<>>><<<<>><<<<>>><<<<>>><<>><<<<><><<<>>>><<>>>><<<>>><<<><<<>>><<<>>><<>>><><<<>>><<>><<>>>><><<<<>>><<<>><<><<>>>><<<><<><<<<>><><><<>>>><<<<>>><<<<>><<><<><>>>><<<>>><<<>>>><<<>>><<<<>>>><<<<>>><<<<><<><<<<>>><>>><<<<><>>><>>>><<>>><<>>>><<<<>>><<<>>><<<<><<>>>><>>><<<>>><<<<><><<>>><<<>>><<<>><<<>>>><<<><<>>>><<>>><>><<<>>>><<>>><<>><<<>>>><<><<<<>>>><<>>>><<<<><<>>><<>><>>>><<<>>><<<><<<><<<>>>><<>>>><<>>>><<<<>>><><<<>>>><<<<>><<>>>><><>><<<>><<<<>><>>>><><<<>>>><>><<<<>>><<<<>><<<>><<>>>><>>><<<>>>><>>>><>>>><<>><<<<>>>><<><<<<>>>><><<><>><<<>>><>>>><<>>><<<>><<<>>>><<><<<<><<<>>>><<<<>><<>><<>><<>>>><>>><<<>>><<<><<<<>>>><>>><>><<>>>><>>>><<<<>><<<>>>><<<>><<<<>>><<>><<><<>>>><>><>>><><<>><<><<><<<>>><<<>><>><<<<>>><<<<>>>><<<<><>>>><<>>>><<<<>><<<<>><<<>>>><>><<<>>><<<<>>>><<<<>>><<<><><<><<<<>><<<<><<>><<<<>>>><>><>>>><<<>><<>><<<<>>><<<<><<<><>><>>>><>><>>><<>>>><<<<>><<>>>><<<<>>><<<>><>><<<<>>><<<<>>><<><<<<>>>><<<<>>><<<>>><>><<<><<>>>><<>><<<<>><<<>><><<<>>><<<>><<<<>>>><><<<>>>><<<<>><<<<>>>><<<<>>><<<><>>><<<<>>><>><>>><<>><<><<<>>>><><<<<>>>><<<<>>><<<<>>><<>><>><<>><<<>>><<<<>>><<>>><<<>>><<<<><>>>><>>><<<<>>><<<><>><<>>><<>><<>>>><>>><<<<>><<>><<<>><><<><>><<<<>>><<<<><<>>><>>>><<>><<<>><<>>><<>>>><<<><<<><<<<>><<<>><<>>><><>><<<><<><>>>><<<>>>><<<>>>><>><<><<<<><<<<>>>><<<>>><<><<<<>>><<<>><<>><>>><<>><<<>>>><><><<<>><>>>><<>>>><>>><<<>><>>>><<<<>><<<<><<<<>><<<<>><<<<>>>><<<<>>><<<><<<<>>>><<>>>><<<>><<>>>><<<>>><<<>>><<>><<<>><<<<><<<>><><>>><<><<>>>><<>>><>>><<<>><<><<><<<>>>><<<<><<<><<<<>><<<<><<<>><<>>><<<<>>>><<<>>>><<><<>>>><><<<>><><>>>><<<<>>><><><<<>>>><<<><<<><<>>><<>><<<>><<<>>>><<<>><<>><<<<>><<>>><><<<<><<><<>><<<>>>><<<><<<>>>><<>><<<<>><<<<><<<>>>><<<<>><<<><<<<>>>><>>>><<<><<><><<<><<<>>><<>>><<<<><<>>>><<<>>><<<<>>><<><<>><<>>>><><<<<>>>><><>>><<<>>><<>>><<<<>>>><<<<>><<<>>>><>>><<>>>><>><<<>><<<><<<<>>><>><<<>><<>>>><<><<<>><<><<<><<<<>>>><><<>>>><<<>>><<<<>><>>>><<>><>>><<<<>>><<<<><<<<>><<<<>><<<><<<<><>>>><><<<>>>><<>><<>>><><<<>>><><<<<><>><<<>>>><<><<<<>>>><<<><<<>>>><>>><<<><<<<>>>><<<<><>>><<<>>><<>><<><<<<>>><<<<>>><><>>><<<<><<>>><<>><<<><>><<<<>>><<>><<<>>><<<<>>><<<>>><<<><<<<><<<<>><<><<<>>><<>><<<>><><<<>>>><<><<>>><<<>>>><<>><<<<>>>><>>><<<<><><<<<>><>>><<<>>><<<>>><<<>><>>>><>>><<<<><<<>><>>>><<<>>><><<><>>><<<>>>><>><<<>><<<<><>><<>>><<<><>>><<<>><><<<<>>>><<<><<<>><>><<<>>><>>>><<<>>><<>>>><<<>>>><>>>><<<<>>><<><<>><<<<>>>><>>><<<><<<<>>>><<<>>><<>>><<<>><<<<>>><<<<><>>>><<>>>><<>>><<>><<<<>><<<>>><<<>>>><<<<>>>><>>><><<>>>><<<>>><<><><<>>>><<<>>>><<>><<<<>><<<<>>><<<><<<><<><<<><<<>>>><<<<>>><<<<>><<<<>><<<<>>>><<<><<>><<>>>><<>>><<<<>>>><<<>>>><><>>>><<<<>>>><<<<>><<<>>><<<<>><<<>>><<<<><<<><<<<>>><<<>>><<>>>><>>>><<<<><<<<>><<>>>><<<<><>>><><<>>><>>>><<<>>>><<<<><<><<<<><><<<<><<<<>>>><<<<>>><<<>>><<<<>>>><<<<>><>>>><<>><<>><<>><<<>><>>><>><>><<><>><<<>>><<<>>><>>><<>><<<><>><<>><<<>><<<>><>>>><>>>><<<>><>>>><<>>>><>><<>>>><>><<<<>><<>>><<<<>>><<>><<<>>>><<>><<>><<<<><<<><<<<>>><>>><<>><>><<<>><<<>>><>>>><<>>>><>>><<>>>><><>><<>><<<<><<>>>><<<><<>>>><<>><<<>>>><>><<>>>><>>><><><<<><>><<><>><<<><<<<>>><<>>>><><<>>>><<>>><<>>>><<<<>>>><>>><><>><>>><><<<<>>>><<<<>>><<<<>>>><<>>>><<>>>><<<>>>><<<<>>>><<<>>><<<>>>><<<<><<<>>><<<<><<>>><<>>>><<<<>><>>>><<><><<<<><<<>>><<<>><<><<<<>>>><<>><><<<>>>><<<>>><<<<>>><>>><<<<>>>><<<<>>><<><<>>>><<<><>>>><<<<>>><>><>>>><<<<>><<>>>><>><<<>><<>>><<<<><<<>>><<<><>>><<<><><<<>>><<<>>>><><>>><<<><<<<>>><>>><>><<<><<<<>>><>>>><<<<>>><<<<>>><<<>><<>>><<<<>>>><<>><<<<>><>>><<<>><>><<<<><<<<>><<<><<<>>>><><>><<>>>><<<<><<<>>>><<<<>><<>>><<<><>>>><<<<><><>>>><<>>><<<>>><<<<>>>><<>>><><<<>>><<>>>><<<<>>><<<<>><><<<<>>>><<<<>><<<><<><<<<>>><<<<><<>><<<>>><<<>>><><>>><<<>>><<>><<<>>><<<<>><>>><>>>><>>><<<<>>><>>><<><<<<>>>><<<<>>><<<<>>>><<<>>><>>><>><<<>>>><>>><<<<>>>><<<<>>><<<<>>>><<<>><<<>><<>><>><<<<>>><>>>><<<<>><<<<><<<<>>><>><>><<<>>>><<<><>>><<><<<<>>>><<<<>><<<<>>>><<<>><<<><<<>>>><<<<><<<>><><<<<>>>><<>>>><<<<>>><<>>><>>><<><<<<><<<><<><<<<>>><>>>><<<<><<<<>>>><<<<>>><<>>>><<>>>><>>>><<><<<<><>><<>>>><<<>>><<<<>><<>>><<<>>>><<<>><<<<>><<<><>>><><<<>>><<<>>><<<><<>><<<>><<<>>>><<>><<>>><<<>><>>>><<>>><<<<><><<<>>><<>>><>><<<<>>><<<>>><<>><<<<><<>>><<>>>><<>><>>><<<<><<<<><<<><<<<>>><<>>><<<<>>><<>>>><>>>><<<<>>>><<<>><<<>><<<><>><>>>><<<><<<>>><>>><<<><<>>>><>>>><>>><><<<>>>><<<>>><>><<>>>><<>>><<>>>><<<<>>>><<>>>><<>>><<<>><<<<>>><><<<>>><<<><<<><<<>>>><>>>><><<>>>><<<>><><<<<>>>><<<<>>><<<<>><>><>>>><<<><><>>><<<>>><<<>>>><<>>>><<<<>>><><<>>>><>>><<><<<>>>><<<>>>><<<>>>><>>><>>>><>>><>>><>>><<<<><>>><<<>>><>><<<>>><>>>><<<<>>><<>>>><<><<>>><<><<<<>>><<>>><<><<<<>>><<>><<>>>><<<>>>><<<><><<<<><<<><<><>><<<<>><<>><<<<>><<>>>><<<>>>><<<<>>>><>>>><<<<>><<<>>>><<>>>><<><<>>><<<<><>>>><<<>>>><<<<><<<<>><>>><<<>>><<<<><>><<<><<>>><<>>>><<<><<<>><<<>>><<<>>><<<<>>><<<<>>>><<<>>>><<<><<<<>><<<<>>><<<>><<<>>><>>><<>><>>>><>>>><<<>>><<<>>>><<>>><<<><><<<>>>><<<><<<<>>><<<>>>><><<>>><<>>><<<>>><<<<>>><<<>>>><<<>>>><>>>><<<><<><>>>><<>>>><<<<>><<<><<<<>>><>",
	}

	pieces := make([][][]string, 0)
	// piece1
	piece1 := make([][]string, 0)
	piece1a := []string{".", ".", "@", "@", "@", "@", "."}
	piece1 = append(piece1, piece1a)

	//piece2
	piece2 := make([][]string, 0)
	piece2c := []string{".", ".", ".", "@", ".", ".", "."}
	piece2b := []string{".", ".", "@", "@", "@", ".", "."}
	piece2a := []string{".", ".", ".", "@", ".", ".", "."}
	piece2 = append(piece2, piece2a)
	piece2 = append(piece2, piece2b)
	piece2 = append(piece2, piece2c)

	//piece3
	piece3 := make([][]string, 0)
	piece3c := []string{".", ".", ".", ".", "@", ".", "."}
	piece3b := []string{".", ".", ".", ".", "@", ".", "."}
	piece3a := []string{".", ".", "@", "@", "@", ".", "."}
	piece3 = append(piece3, piece3a)
	piece3 = append(piece3, piece3b)
	piece3 = append(piece3, piece3c)

	//piece4
	piece4 := make([][]string, 0)
	piece4d := []string{".", ".", "@", ".", ".", ".", "."}
	piece4c := []string{".", ".", "@", ".", ".", ".", "."}
	piece4b := []string{".", ".", "@", ".", ".", ".", "."}
	piece4a := []string{".", ".", "@", ".", ".", ".", "."}
	piece4 = append(piece4, piece4a)
	piece4 = append(piece4, piece4b)
	piece4 = append(piece4, piece4c)
	piece4 = append(piece4, piece4d)

	//piece5
	piece5 := make([][]string, 0)
	piece5b := []string{".", ".", "@", "@", ".", ".", "."}
	piece5a := []string{".", ".", "@", "@", ".", ".", "."}
	piece5 = append(piece5, piece5a)
	piece5 = append(piece5, piece5b)

	//append pieces
	pieces = append(pieces, piece1)
	pieces = append(pieces, piece2)
	pieces = append(pieces, piece3)
	pieces = append(pieces, piece4)
	pieces = append(pieces, piece5)

	p := &piece{
		index:  0,
		pieces: pieces,
	}

	arr := make([][]string, 0)
	cache := make(map[MyKey]MyValue)
	b := &board{
		columns: arr,
		_cache:  cache,
	}
	b.appendEmptyLine()
	b.appendEmptyLine()
	b.appendEmptyLine()
	b.appendEmptyLine()
	fmt.Println("Before preparation")
	b.print()
	b.prepareBoardForPieceStart()
	fmt.Println("After preparation")
	b.print()

	for i := 0; i < TOTAL_PIECES; i++ {
		fmt.Printf("Piece %d\n", i)
		b.prepareBoardForPieceStart()
		// fmt.Println("After preparation")
		// b.print()
		b.putPieceInBoard(p.next())
		// fmt.Println("After put piece")
		// b.print()
		// fmt.Println("Moving piece laterally")
		b.movePieceLaterally(d.next())
		for {
			// fmt.Println("Moving piece down")
			if b.canPieceMoveDown() {
				b.movePieceDown()
			} else {
				break
			}

			// fmt.Println("Moving piece laterally")
			b.movePieceLaterally(d.next())
			// fmt.Println("After moving piece laterally")
			// b.print()
		}
		b.settlePiece()
		// fmt.Println("After settling piece")
		// b.print()
		if b.highestLevel() > 50 && b.foundLoop == false {
			found, currentHeight, lastHeigth, pieceCount, lastPieceCount := b._check_cache(p.index, d.index, b.getStringRepresentationTopXRows(50), i+1)
			if found {
				b.foundLoop = true
				fmt.Printf("Found a loop currentHeight %d lastHeight %d currentPiece %d lastPiece %d\n", currentHeight, lastHeigth, pieceCount, lastPieceCount)
				heightDelta := currentHeight - lastHeigth
				piecesDelta := pieceCount - lastPieceCount
				fmt.Printf("heightDelta %d piecesDelta %d\n", heightDelta, piecesDelta)
				repeatAddedHeight, remainingPiecesAfterRepetition := b.calculateHeight(TOTAL_PIECES, i+1, piecesDelta, heightDelta)
				fmt.Printf("repeatAddedHeight %d remainingPiecesAfterRepetition %d\n", repeatAddedHeight, remainingPiecesAfterRepetition)
				b.repeatsHeight = repeatAddedHeight
				i = TOTAL_PIECES - remainingPiecesAfterRepetition - 1
			}
		}
	}
	fmt.Printf("Response is %d\n", b.highestLevel()+1+b.repeatsHeight)
}

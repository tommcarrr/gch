package board

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var fileMap = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
}

var rankMap = map[string]int{
	"1": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
}

const errInvalidFen string = "invalid FEN"

const newGameFen string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
const fenPieces string = "rnbqkpPRNBQK"
const fenNums string = "12345678"

type BoardDefinition struct {
	pieces         []string
	toPlay         string
	castlingRights string
	enPasseMove    string
	halfMoveCount  int
	fullMoveCount  int
}

func (b BoardDefinition) piecesString() (str string) {
	return strings.Join(b.pieces, "")
}

func (b BoardDefinition) String() string {

	s := fmt.Sprintf("%s to play\n\n", b.toPlay)

	for i, square := range b.pieces {
		s = s + fmt.Sprint(square) + " "
		if (i+1)%8 == 0 {
			s = s + "\n"
		}
	}

	return s
}

func NewGame() (b BoardDefinition) {
	b, _ = generateBoardFromFen(newGameFen)
	return
}

func NewGameFromFen(fen string) (b BoardDefinition, err error) {
	b, err = generateBoardFromFen(fen)
	return
}

func generateBoardFromFen(fen string) (b BoardDefinition, err error) {

	elements := strings.Split(fen, " ")

	if len(elements) != 6 {
		err = errors.New(errInvalidFen)
		return
	}

	piecesSegment := elements[0]
	b.pieces = parseFenPieces(piecesSegment)

	if len(b.pieces) != 64 {
		err = errors.New(errInvalidFen)
		return
	}

	b.toPlay = elements[1]
	b.castlingRights = elements[2]
	b.enPasseMove = elements[3]
	b.halfMoveCount, _ = strconv.Atoi(elements[4])
	b.fullMoveCount, _ = strconv.Atoi(elements[5])

	return
}

func parseFenPieces(piecesSegment string) (pieces []string) {
	piecesString := strings.Split(piecesSegment, "")

	for _, val := range piecesString {
		if strings.Contains(fenPieces, val) {
			pieces = append(pieces, val)
		} else if strings.Contains(fenNums, val) {
			fenNum, _ := strconv.Atoi(val)
			for i := 0; i < fenNum; i++ {
				pieces = append(pieces, "-")
			}
		}
	}
	return
}

func getPieceFromSquare(b BoardDefinition, square string) (piece string, err error) {
	if len(square) != 2 {
		err = errors.New("invalid square")
	}

	squareRunes := []rune(square)

	file := fileMap[string(squareRunes[0])]
	rank := rankMap[string(squareRunes[1])]

	index := file + (8 * (7 - rank))
	piece = b.pieces[index]

	return
}

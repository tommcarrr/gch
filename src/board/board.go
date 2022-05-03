package board

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

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

func (b *BoardDefinition) movePieceFromString(move string) (err error) {

	if len(move) < 4 || len(move) > 5 {
		err = errors.New("invalid move format")
		return
	}

	moveRunes := []rune(move)

	from := string(moveRunes[0:2])
	to := string(moveRunes[2:4])

	promotion := ""
	if len(move) == 5 {
		promotion = string(moveRunes[5])
	}

	fromIndex := squareMap[from]
	toIndex := squareMap[to]

	fromPiece := b.pieces[fromIndex]
	toPiece := b.pieces[toIndex]

	if toPiece != "-" && IsUpper(fromPiece) == IsUpper(toPiece) {
		err = errors.New("invalid move")
		return
	} else {
		b.pieces[fromIndex] = "-"
		if promotion != "" {
			b.pieces[toIndex] = promotion
		} else {
			b.pieces[toIndex] = fromPiece
		}
	}
	return
}

func IsUpper(s string) bool {
	return s == strings.ToUpper(s)
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

func (b BoardDefinition) getPieceFromSquare(square string) (piece string, err error) {

	if len(square) != 2 {
		err = errors.New("invalid square")
	}

	index := squareMap[square]
	piece = b.pieces[index]

	return
}

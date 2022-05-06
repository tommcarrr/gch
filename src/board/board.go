package board

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const newGameFen string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
const fenPieces string = "rnbqkpPRNBQK"
const fenNums string = "12345678"

type BoardDefinition struct {
	pieces         []string
	toPlay         Color
	castlingRights string
	enPasseMove    string
	halfMoveCount  int
	fullMoveCount  int
}

func (b *BoardDefinition) flip() {
	if b.toPlay == White {
		b.toPlay = Black
	} else {
		b.toPlay = White
	}
}

func (b BoardDefinition) piecesString() (str string) {
	return strings.Join(b.pieces, "")
}

func (b BoardDefinition) String() string {

	s := fmt.Sprintf("\u001b[2J%v to play\n\n", b.toPlay)

	cnt := 0

	for i, square := range b.pieces {
		if cnt%2 == 0 {
			s = s + "\u001b[47m"
		} else {
			s = s + "\u001b[40m"
		}
		s = s + pieceMap[square] + " "
		if (i+1)%8 == 0 {
			s = s + "\u001b[0m\n"
			cnt++
		}
		cnt++
	}

	s = s + "\u001b[0m"

	return s
}

func (b *BoardDefinition) MovePieceFromString(move string) (err error) {

	if len(move) < 4 || len(move) > 5 {
		err = errors.New(errInvalidMoveFormat)
		return
	}

	moveRunes := []rune(move)

	from := string(moveRunes[0:2])
	to := string(moveRunes[2:4])

	promotion := ""
	if len(move) == 5 {
		promotion = string(moveRunes[4])
		if b.toPlay == White {
			promotion = strings.ToUpper(promotion)
		} else {
			promotion = strings.ToLower(promotion)
		}
	}

	fromIndex := squareMap[from]
	toIndex := squareMap[to]

	fromPiece := b.pieces[fromIndex]
	toPiece := b.pieces[toIndex]

	if fromPiece == "-" {
		err = errors.New(errInvalidMoveNoPiece)
		return
	}

	if (b.toPlay == White && !IsUpper(fromPiece)) || (b.toPlay == Black && IsUpper(fromPiece)) {
		err = errors.New(errInvalidMoveWrongColor)
		return
	}

	if toPiece != "-" && IsUpper(fromPiece) == IsUpper(toPiece) {
		err = errors.New(errInvalidMove)
		return
	} else {
		b.pieces[fromIndex] = "-"
		if promotion != "" {
			b.pieces[toIndex] = promotion
		} else {
			b.pieces[toIndex] = fromPiece
		}
	}

	b.flip()

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

	if elements[1] == "w" {
		b.toPlay = White
	} else if elements[1] == "b" {
		b.toPlay = Black
	} else {
		err = errors.New(errInvalidFen)
		return
	}
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
		err = errors.New(errInvalidSquare)
	}

	index := squareMap[square]
	piece = b.pieces[index]

	return
}

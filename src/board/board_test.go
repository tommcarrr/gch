package board

import (
	"testing"
)

func TestNewBoard(t *testing.T) {

	board := NewGame()

	got := board.piecesString()
	want := "rnbqkbnrpppppppp--------------------------------PPPPPPPPRNBQKBNR"
	if got != want {
		t.Errorf("board.pieces: got %v, wanted %v", got, want)
	}

	got = board.toPlay
	want = "w"
	if got != want {
		t.Errorf("board.toPlay: got %v, wanted %v", got, want)
	}

	got = board.enPasseMove
	want = "-"
	if got != want {
		t.Errorf("enPasseMove.toPlay: got %v, wanted %v", got, want)
	}

	got = board.castlingRights
	want = "KQkq"
	if got != want {
		t.Errorf("enPasseMove.castlingRights: got %v, wanted %v", got, want)
	}

	goti := board.halfMoveCount
	wanti := 0
	if goti != wanti {
		t.Errorf("enPasseMove.halfMoveCount: got %v, wanted %v", goti, wanti)
	}

	goti = board.fullMoveCount
	wanti = 1
	if goti != wanti {
		t.Errorf("enPasseMove.fullMoveCount: got %v, wanted %v", goti, wanti)
	}
}

func TestNewBoardFromFen(t *testing.T) {

	fen := "rnbqkbnr/ppppp3/7p/6p1/3P4/4P1p1/PPP2PPP/RN1QKBNR w KQkq - 0 6"
	board, _ := NewGameFromFen(fen)

	got := board.piecesString()
	want := "rnbqkbnrppppp----------p------p----P--------P-p-PPP--PPPRN-QKBNR"
	if got != want {
		t.Errorf("board.pieces: got %v, wanted %v", got, want)
	}

	got = board.toPlay
	want = "w"
	if got != want {
		t.Errorf("board.toPlay: got %v, wanted %v", got, want)
	}

	got = board.enPasseMove
	want = "-"
	if got != want {
		t.Errorf("enPasseMove.toPlay: got %v, wanted %v", got, want)
	}

	got = board.castlingRights
	want = "KQkq"
	if got != want {
		t.Errorf("enPasseMove.castlingRights: got %v, wanted %v", got, want)
	}

	goti := board.halfMoveCount
	wanti := 0
	if goti != wanti {
		t.Errorf("enPasseMove.halfMoveCount: got %v, wanted %v", goti, wanti)
	}

	goti = board.fullMoveCount
	wanti = 6
	if goti != wanti {
		t.Errorf("enPasseMove.fullMoveCount: got %v, wanted %v", goti, wanti)
	}
}

func TestNewBoardFromFen2(t *testing.T) {

	fen := "r1bk3r/p2p1pNp/n2B1n2/1p1NP2P/6P1/3P4/P1P1K3/q5b1 w - - 0 23"
	board, _ := NewGameFromFen(fen)

	got := board.piecesString()
	want := "r-bk---rp--p-pNpn--B-n---p-NP--P------P----P----P-P-K---q-----b-"
	if got != want {
		t.Errorf("board.pieces: got %v, wanted %v", got, want)
	}

	got = board.toPlay
	want = "w"
	if got != want {
		t.Errorf("board.toPlay: got %v, wanted %v", got, want)
	}

	got = board.enPasseMove
	want = "-"
	if got != want {
		t.Errorf("enPasseMove.toPlay: got %v, wanted %v", got, want)
	}

	got = board.castlingRights
	want = "-"
	if got != want {
		t.Errorf("enPasseMove.castlingRights: got %v, wanted %v", got, want)
	}

	goti := board.halfMoveCount
	wanti := 0
	if goti != wanti {
		t.Errorf("enPasseMove.halfMoveCount: got %v, wanted %v", goti, wanti)
	}

	goti = board.fullMoveCount
	wanti = 23
	if goti != wanti {
		t.Errorf("enPasseMove.fullMoveCount: got %v, wanted %v", goti, wanti)
	}
}

func TestNewBoardFromFenInvalidCount(t *testing.T) {

	fen := "r1bk3r/p2p1pNp/n2B1n2/1p1NP2P/6P1/3P4/P1P1K3/q5b1 w - 0 23"
	_, err := NewGameFromFen(fen)

	got := err.Error()
	want := errInvalidFen

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestNewBoardFromFenInvalidPieces(t *testing.T) {

	fen := "r1bk3r/p2p1pNp/n2B1n2/1p1NP2P/6P1/3P4/P1P1K3/q5 w - - 0 23"
	_, err := NewGameFromFen(fen)

	got := err.Error()
	want := errInvalidFen

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetPieceA1(t *testing.T) {
	square := "a1"
	board := NewGame()

	want := "R"
	got, _ := board.getPieceFromSquare(square)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetPieceE4(t *testing.T) {
	square := "e4"
	board := NewGame()

	want := "-"
	got, _ := board.getPieceFromSquare(square)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetPieceH8(t *testing.T) {
	square := "h8"
	board := NewGame()

	want := "r"
	got, _ := board.getPieceFromSquare(square)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetPieceA1Fen(t *testing.T) {
	square := "a1"
	fen := "r1bk3r/p2p1pNp/n2B1n2/1p1NP2P/6P1/3P4/P1P1K3/q5b1 w - - 0 23"
	board, _ := NewGameFromFen(fen)

	want := "q"
	got, _ := board.getPieceFromSquare(square)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetPieceD6Fen(t *testing.T) {
	square := "d6"
	fen := "r1bk3r/p2p1pNp/n2B1n2/1p1NP2P/6P1/3P4/P1P1K3/q5b1 w - - 0 23"
	board, _ := NewGameFromFen(fen)

	want := "B"
	got, _ := board.getPieceFromSquare(square)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMoveD4(t *testing.T) {

	move := "d2d4"
	board := NewGame()

	board.movePieceFromString(move)

	square := "d4"
	want := "P"
	got, _ := board.getPieceFromSquare(square)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	square = "d2"

	want = "-"
	got, _ = board.getPieceFromSquare(square)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMoveE4E5(t *testing.T) {

	move1 := "e2e4"
	move2 := "e7e5"
	board := NewGame()

	board.movePieceFromString(move1)
	board.movePieceFromString(move2)

	square := "e4"
	want := "P"
	got, _ := board.getPieceFromSquare(square)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	square = "e2"

	want = "-"
	got, _ = board.getPieceFromSquare(square)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	square = "e5"

	want = "p"
	got, _ = board.getPieceFromSquare(square)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	square = "e7"

	want = "-"
	got, _ = board.getPieceFromSquare(square)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMoveInvalidFormat(t *testing.T) {
	move := "Pd4"
	board := NewGame()

	err := board.movePieceFromString(move)

	want := errInvalidMoveFormat
	got := err.Error()

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMoveInvalid(t *testing.T) {
	move := "b1d2"
	board := NewGame()

	err := board.movePieceFromString(move)

	want := errInvalidMove
	got := err.Error()

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

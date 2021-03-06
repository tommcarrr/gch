package board

type Color int8

const (
	Black Color = -1
	White Color = 1
)

func (c Color) String() string {
	switch c {
	case Black:
		return "Black"
	case White:
		return "White"
	default:
		return "Color Not Found"
	}
}

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

var squareMap = map[string]int{
	"a8": 0,
	"b8": 1,
	"c8": 2,
	"d8": 3,
	"e8": 4,
	"f8": 5,
	"g8": 6,
	"h8": 7,
	"a7": 8,
	"b7": 9,
	"c7": 10,
	"d7": 11,
	"e7": 12,
	"f7": 13,
	"g7": 14,
	"h7": 15,
	"a6": 16,
	"b6": 17,
	"c6": 18,
	"d6": 19,
	"e6": 20,
	"f6": 21,
	"g6": 22,
	"h6": 23,
	"a5": 24,
	"b5": 25,
	"c5": 26,
	"d5": 27,
	"e5": 28,
	"f5": 29,
	"g5": 30,
	"h5": 31,
	"a4": 32,
	"b4": 33,
	"c4": 34,
	"d4": 35,
	"e4": 36,
	"f4": 37,
	"g4": 38,
	"h4": 39,
	"a3": 40,
	"b3": 41,
	"c3": 42,
	"d3": 43,
	"e3": 44,
	"f3": 45,
	"g3": 46,
	"h3": 47,
	"a2": 48,
	"b2": 49,
	"c2": 50,
	"d2": 51,
	"e2": 52,
	"f2": 53,
	"g2": 54,
	"h2": 55,
	"a1": 56,
	"b1": 57,
	"c1": 58,
	"d1": 59,
	"e1": 60,
	"f1": 61,
	"g1": 62,
	"h1": 63,
}

var pieceMap = map[string]string{
	"K": "\u001b[37m\u2654",
	"Q": "\u001b[37m\u2655",
	"R": "\u001b[37m\u2656",
	"B": "\u001b[37m\u2657",
	"N": "\u001b[37m\u2658",
	"P": "\u001b[37m\u2659",
	"k": "\u001b[30;1m\u265a",
	"q": "\u001b[30;1m\u265b",
	"r": "\u001b[30;1m\u265c",
	"b": "\u001b[30;1m\u265d",
	"n": "\u001b[30;1m\u265e",
	"p": "\u001b[30;1m\u265f",
	"-": " ",
}

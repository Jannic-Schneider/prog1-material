package chesspiece

type PType int

const (
	BISHOP PType = iota
	QUEEN
	KNIGHT
	ROOK
	PAWN
)

type Colour int

const (
	WHITE Colour = iota
	Black
)

type ChessPiece struct {
	pieceType PType
	colour    Colour
	row       int
	column    int
}

func (p ChessPiece) MoveAllowed(row, col int) bool {
	if row < 0 || row > 7 || col < 0 || col > 7 {
		return false
	}
	if p.pieceType == BISHOP {
		if row-col == p.row-p.column {
			return true
		}
		if row+col == p.row+p.column {
			return true
		}
	}

	if p.pieceType == KNIGHT {
		if row+2 == p.row && col+1 == p.column {
			return true
		}
		if row+2 == p.row && col-1 == p.column {
			return true
		}
		if row-2 == p.row && col+1 == p.column {
			return true
		}
		if row-2 == p.row && col-1 == p.column {
			return true
		}
		if row+1 == p.row && col+2 == p.column {
			return true
		}
		if row+1 == p.row && col-2 == p.column {
			return true
		}
		if row-1 == p.row && col+2 == p.column {
			return true
		}
		if row-1 == p.row && col-2 == p.column {
			return true
		}
	}
	if p.pieceType == ROOK {
		if row == p.row || col == p.column {
			return true

		}
	}
	if p.pieceType == QUEEN {
		if row == p.row || col == p.column {
			return true
		}
		if row-col == p.row-p.column {
			return true
		}
		if row+col == p.row+p.column {
			return true
		}

	}
	return false
}

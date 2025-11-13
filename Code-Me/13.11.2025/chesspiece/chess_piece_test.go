package chesspiece

import "fmt"

func ExampleChessPiece_MoveAllowed_rook() {
	l1 := ChessPiece{
		pieceType: BISHOP,
		colour:    WHITE,
		row:       2,
		column:    2,
	}
	k1 := ChessPiece{
		pieceType: KNIGHT,
		colour:    WHITE,
		row:       2,
		column:    2,
	}

		k2 := ChessPiece{
		pieceType: KNIGHT,
		colour:    WHITE,
		row:       0,
		column:    0,
	}

		r1 := ChessPiece{
		pieceType: ROOK,
		colour:    WHITE,
		row:       2,
		column:    2,
	}
		
	    q1 := ChessPiece{
		pieceType: QUEEN,
		colour:    WHITE,
		row:       2,
		column:    2,
	}


	fmt.Println(l1.MoveAllowed(3, 4))
	fmt.Println(l1.MoveAllowed(0, 0))
	fmt.Println(l1.MoveAllowed(9, 9))

	fmt.Println(k1.MoveAllowed(4, 3))
	fmt.Println(k1.MoveAllowed(0, 0))

	fmt.Println(k2.MoveAllowed(-2, -1))

    fmt.Println(r1.MoveAllowed(5,2 ))
	fmt.Println(r1.MoveAllowed(4,3 ))

	fmt.Println(q1.MoveAllowed(4,2 ))
	fmt.Println(q1.MoveAllowed(4,3 ))
	fmt.Println(q1.MoveAllowed(5,5 ))

	//Output:
	//false
	//true
	//false
	//true
	//false
	//false
	//true
	//false
	//true
	//false
	//true
}

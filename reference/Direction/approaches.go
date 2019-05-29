package direction

// Movimento di 1 casella verso x
func LowSpeed_Mov(direction string){
	MoveInDirection(direction, 10)
}

// Movimento di 3 casella verso x
func MidSpeed_Mov(direction string){
	MoveInDirection(direction, 15)
}

// Movimento di 6 caselle  (libero)
func MaxSpeed_Mov(direction string){
	MoveInDirection(direction, 30)
}
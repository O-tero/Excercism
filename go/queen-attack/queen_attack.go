package queenattack

import (
	"errors"
)

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if len(whitePosition) != 2 || len(blackPosition) != 2 {
		return false, errors.New("invalid position format")
	}
	if whitePosition == blackPosition {
		return false, errors.New("queens cannot occupy the same position")
	}

	// Convert chess notation to numerical coordinates
	col1, row1 := int(whitePosition[0]-'a'), int(whitePosition[1]-'1')
	col2, row2 := int(blackPosition[0]-'a'), int(blackPosition[1]-'1')

	// Ensure coordinates are within valid range (0 to 7)
	if col1 < 0 || col1 > 7 || row1 < 0 || row1 > 7 || col2 < 0 || col2 > 7 || row2 < 0 || row2 > 7 {
		return false, errors.New("position out of bounds")
	}

	// Check if queens can attack each other
	if row1 == row2 || col1 == col2 || abs(row1-row2) == abs(col1-col2) {
		return true, nil
	}
	
	return false, nil
}

// Helper function to compute absolute difference
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
